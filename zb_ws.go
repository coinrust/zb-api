package zb

import (
	"context"
	"fmt"
	"github.com/recws-org/recws"
	"github.com/tidwall/gjson"
	"strings"

	//"github.com/tidwall/gjson"
	"log"
	"time"
)

const (
	wsURL = "wss://api.zb.live/websocket"
)

type ZBWebsocket struct {
	wsURL     string
	accessKey string
	secretKey string
	debugMode bool

	ctx    context.Context
	cancel context.CancelFunc
	conn   recws.RecConn

	subscriptions  map[string]interface{}
	tickerCallback func(ticker *WSTicker)
	depthCallback  func(depth *WSDepth)
	tradesCallback func(trades *WSTrades)
}

func (ws *ZBWebsocket) Start() {
	log.Printf("wsURL: %v", ws.wsURL)
	ws.conn.Dial(ws.wsURL, nil)
	go ws.run()
}

func (ws *ZBWebsocket) SetTickerCallback(callback func(ticker *WSTicker)) {
	ws.tickerCallback = callback
}

func (ws *ZBWebsocket) SetDepthCallback(callback func(depth *WSDepth)) {
	ws.depthCallback = callback
}

func (ws *ZBWebsocket) SetTradesCallback(callback func(trades *WSTrades)) {
	ws.tradesCallback = callback
}

func (ws *ZBWebsocket) sendWSMessage(msg interface{}) error {
	return ws.conn.WriteJSON(msg)
}

func (ws *ZBWebsocket) SubscribeTicker(symbol string) (err error) {
	ch := fmt.Sprintf("%v_ticker", symbol) // zbqc_ticker
	return ws.Subscribe("addChannel", ch)
}

func (ws *ZBWebsocket) SubscribeDepth(symbol string) (err error) {
	ch := fmt.Sprintf("%v_depth", symbol) // zbqc_depth
	return ws.Subscribe("addChannel", ch)
}

func (ws *ZBWebsocket) SubscribeTrades(symbol string) (err error) {
	ch := fmt.Sprintf("%v_trades", symbol) // zbqc_trades
	return ws.Subscribe("addChannel", ch)
}

func (ws *ZBWebsocket) Subscribe(event string, channel string) (err error) {
	params := map[string]string{
		"event":   event,
		"channel": channel,
	}
	id := event + ":" + channel
	ws.subscriptions[id] = params
	err = ws.sendWSMessage(params)
	return
}

func (ws *ZBWebsocket) run() {
	ctx := context.Background()
	for {
		select {
		case <-ctx.Done():
			go ws.conn.Close()
			log.Printf("Websocket closed %s", ws.conn.GetURL())
			return
		default:
			messageType, msg, err := ws.conn.ReadMessage()
			if err != nil {
				log.Printf("Read error: %v", err)
				time.Sleep(100 * time.Millisecond)
				continue
			}

			ws.handleMsg(messageType, msg)
		}
	}
}

func (ws *ZBWebsocket) handleMsg(messageType int, msg []byte) (err error) {
	ret := gjson.ParseBytes(msg)
	// {"no":"0","code":1008,"success":false,"message":"事件为空"}
	// {"date":"1587976888747","ticker":{"high":"1.755","vol":"12321359.11","last":"1.7378","low":"1.7203","buy":"1.7338","sell":"1.74"},"dataType":"ticker","channel":"zbqc_ticker"}
	// {"asks":[[7.1557,24.6000],[7.1556,24.6000],[7.1555,1066.7624],[7.1554,251.2788],[7.1553,198553.8390],[7.1552,201429.2054],[7.1551,198714.6994],[7.155,206158.4851],[7.1549,198490.2279],[7.1548,200346.2849],[7.1547,10291.6508],[7.1546,35.1323],[7.1545,111.7031],[7.1544,65.3570],[7.1543,2088.0385],[7.1542,122.0128],[7.1541,64.5330],[7.154,229.8275],[7.1539,2086.0695],[7.1538,41378.8821],[7.1537,2109.0169],[7.1536,49.6677],[7.1535,49.6677],[7.1534,49.6432],[7.1533,59.5139],[7.1532,44.6713],[7.1531,44.6491],[7.153,18787.0153],[7.1529,1039.4362],[7.1528,46.4324],[7.1527,54.1880],[7.1526,745.9965],[7.1525,58.6652],[7.1524,44.5297],[7.1523,159.1413],[7.1522,44.5297],[7.1521,44.4456],[7.152,224.4848],[7.1519,57.9773],[7.1518,503.5241],[7.1517,44.3497],[7.1516,53.1305],[7.1515,707.3196],[7.1514,1950.4191],[7.1513,1134.5073],[7.1512,388.4986],[7.1511,19.5226],[7.151,332.0070],[7.1509,202.2508],[7.1508,219.0730]],"dataType":"depth","bids":[[7.1507,142.3197],[7.1506,56.2844],[7.1505,44.2844],[7.1504,589.6786],[7.1503,1684.4568],[7.1502,48.8675],[7.1501,1781.3638],[7.15,16856.2904],[7.1499,51.9595],[7.1498,2900.2164],[7.1497,45.2164],[7.1496,44.2164],[7.1495,28.9854],[7.1494,632.8461],[7.1493,48.0375],[7.1492,43.9222],[7.1491,43.9222],[7.149,204.9710],[7.1489,44.1202],[7.1488,44.1202],[7.1487,2078.6377],[7.1486,56.3954],[7.1485,459.1042],[7.1484,36.8752],[7.1483,52.8039],[7.1482,52.1302],[7.1481,56.1103],[7.148,235.7559],[7.1479,198464.1777],[7.1478,198475.9069],[7.1477,198467.8303],[7.1476,198477.9349],[7.1475,198478.0447],[7.1474,198494.1322],[7.1473,51.9065],[7.1472,47.8063],[7.1471,47.6905],[7.147,231.6876],[7.1469,55.7601],[7.1468,59.6022],[7.1467,47.8714],[7.1466,55.0648],[7.1465,55.0488],[7.1464,703.6178],[7.1463,47.2257],[7.1462,47.2257],[7.1461,54.5094],[7.146,48985.5596],[7.1459,47.1105],[7.1458,28.2627]],"channel":"usdtqc_depth","timestamp":1587977822}
	// {"data":[{"date":1588059873,"amount":"0.55","price":"1.7278","trade_type":"bid","type":"buy","tid":607502815},{"date":1588059875,"amount":"0.70","price":"1.724","trade_type":"bid","type":"buy","tid":607502817},{"date":1588059875,"amount":"0.12","price":"1.7213","trade_type":"bid","type":"buy","tid":607502818},{"date":1588059875,"amount":"2.00","price":"1.7244","trade_type":"bid","type":"buy","tid":607502819},{"date":1588059875,"amount":"0.26","price":"1.7213","trade_type":"bid","type":"buy","tid":607502820},{"date":1588059875,"amount":"0.44","price":"1.7278","trade_type":"bid","type":"buy","tid":607502821},{"date":1588059877,"amount":"943.41","price":"1.7238","trade_type":"bid","type":"buy","tid":607502823},{"date":1588059877,"amount":"4.10","price":"1.723","trade_type":"ask","type":"sell","tid":607502824},{"date":1588059877,"amount":"3.52","price":"1.7262","trade_type":"ask","type":"sell","tid":607502825},{"date":1588059877,"amount":"1.41","price":"1.7226","trade_type":"ask","type":"sell","tid":607502826},{"date":1588059877,"amount":"1.80","price":"1.7277","trade_type":"ask","type":"sell","tid":607502827},{"date":1588059877,"amount":"0.27","price":"1.7251","trade_type":"ask","type":"sell","tid":607502828},{"date":1588059877,"amount":"5.18","price":"1.7217","trade_type":"ask","type":"sell","tid":607502829},{"date":1588059877,"amount":"0.11","price":"1.7271","trade_type":"ask","type":"sell","tid":607502830},{"date":1588059877,"amount":"4.19","price":"1.7226","trade_type":"ask","type":"sell","tid":607502831},{"date":1588059877,"amount":"2.94","price":"1.7201","trade_type":"ask","type":"sell","tid":607502832},{"date":1588059881,"amount":"0.05","price":"1.723","trade_type":"bid","type":"buy","tid":607502833},{"date":1588059881,"amount":"0.01","price":"1.7277","trade_type":"bid","type":"buy","tid":607502834},{"date":1588059881,"amount":"0.01","price":"1.7246","trade_type":"bid","type":"buy","tid":607502835},{"date":1588059881,"amount":"0.04","price":"1.7263","trade_type":"bid","type":"buy","tid":607502836},{"date":1588059881,"amount":"0.02","price":"1.7266","trade_type":"bid","type":"buy","tid":607502837},{"date":1588059881,"amount":"0.11","price":"1.724","trade_type":"bid","type":"buy","tid":607502838},{"date":1588059881,"amount":"0.17","price":"1.7221","trade_type":"bid","type":"buy","tid":607502839},{"date":1588059881,"amount":"0.06","price":"1.7278","trade_type":"bid","type":"buy","tid":607502840},{"date":1588059883,"amount":"0.45","price":"1.7214","trade_type":"bid","type":"buy","tid":607502842},{"date":1588059883,"amount":"0.31","price":"1.7202","trade_type":"bid","type":"buy","tid":607502843},{"date":1588059883,"amount":"0.18","price":"1.7209","trade_type":"bid","type":"buy","tid":607502844},{"date":1588059883,"amount":"0.46","price":"1.7251","trade_type":"bid","type":"buy","tid":607502845},{"date":1588059883,"amount":"0.20","price":"1.7278","trade_type":"bid","type":"buy","tid":607502846},{"date":1588059884,"amount":"1073.48","price":"1.7239","trade_type":"bid","type":"buy","tid":607502848},{"date":1588059886,"amount":"0.08","price":"1.7252","trade_type":"bid","type":"buy","tid":607502849},{"date":1588059886,"amount":"0.02","price":"1.7221","trade_type":"bid","type":"buy","tid":607502850},{"date":1588059886,"amount":"0.11","price":"1.7225","trade_type":"bid","type":"buy","tid":607502851},{"date":1588059886,"amount":"0.02","price":"1.7223","trade_type":"bid","type":"buy","tid":607502852},{"date":1588059886,"amount":"0.06","price":"1.726","trade_type":"bid","type":"buy","tid":607502853},{"date":1588059886,"amount":"0.11","price":"1.7205","trade_type":"bid","type":"buy","tid":607502854},{"date":1588059886,"amount":"0.06","price":"1.7228","trade_type":"bid","type":"buy","tid":607502855},{"date":1588059886,"amount":"0.03","price":"1.7213","trade_type":"bid","type":"buy","tid":607502856},{"date":1588059886,"amount":"0.07","price":"1.7278","trade_type":"bid","type":"buy","tid":607502857},{"date":1588059889,"amount":"10.58","price":"1.7205","trade_type":"ask","type":"sell","tid":607502859},{"date":1588059889,"amount":"8.79","price":"1.725","trade_type":"ask","type":"sell","tid":607502860},{"date":1588059889,"amount":"9.83","price":"1.7218","trade_type":"ask","type":"sell","tid":607502861},{"date":1588059889,"amount":"38.95","price":"1.7222","trade_type":"ask","type":"sell","tid":607502862},{"date":1588059889,"amount":"26.77","price":"1.721","trade_type":"ask","type":"sell","tid":607502863},{"date":1588059889,"amount":"13.56","price":"1.7201","trade_type":"ask","type":"sell","tid":607502864},{"date":1588059890,"amount":"0.09","price":"1.7225","trade_type":"bid","type":"buy","tid":607502865},{"date":1588059890,"amount":"0.21","price":"1.7272","trade_type":"bid","type":"buy","tid":607502866},{"date":1588059890,"amount":"0.13","price":"1.7214","trade_type":"bid","type":"buy","tid":607502867},{"date":1588059890,"amount":"0.27","price":"1.7265","trade_type":"bid","type":"buy","tid":607502868},{"date":1588059890,"amount":"0.10","price":"1.7278","trade_type":"bid","type":"buy","tid":607502869}],"dataType":"trades","channel":"zbqc_trades"}
	log.Printf("%v", string(msg))
	if channelValue := ret.Get("channel"); channelValue.Exists() {
		channel := channelValue.String()
		if strings.HasSuffix(channel, "_ticker") { // zbqc_ticker
			var data WSTicker
			err = json.Unmarshal(msg, &data)
			if err != nil {
				return
			}
			if ws.tickerCallback != nil {
				ws.tickerCallback(&data)
			}
		} else if strings.HasSuffix(channel, "_depth") { // zbqc_depth
			var data WSDepth
			err = json.Unmarshal(msg, &data)
			if err != nil {
				return
			}
			if ws.depthCallback != nil {
				ws.depthCallback(&data)
			}
		} else if strings.HasSuffix(channel, "_trades") { // zbqc_trades
			var data WSTrades
			err = json.Unmarshal(msg, &data)
			if err != nil {
				return
			}
			if ws.tradesCallback != nil {
				ws.tradesCallback(&data)
			}
		}
	}
	return nil
}

func (ws *ZBWebsocket) subscribeHandler() error {
	log.Printf("subscribeHandler")
	for _, v := range ws.subscriptions {
		//log.Printf("sub: %#v", v)
		err := ws.sendWSMessage(v)
		if err != nil {
			log.Printf("%v", err)
		}
	}
	return nil
}

func NewZBWebsocket(accessKey string, secretKey string, debugMode bool) *ZBWebsocket {
	ws := &ZBWebsocket{
		wsURL:         wsURL,
		subscriptions: make(map[string]interface{}),
	}
	ws.ctx, ws.cancel = context.WithCancel(context.Background())
	ws.conn = recws.RecConn{
		KeepAliveTimeout: 10 * time.Second,
	}
	ws.conn.SubscribeHandler = ws.subscribeHandler
	return ws
}
