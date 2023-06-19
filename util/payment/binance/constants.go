package binance

import "time"

const BINANCE_REQUEST_STATUS_SUCCESS = "SUCCESS"
const BINANCE_REQUEST_STATUS_FAIL = "FAIL"
const ORDER_EXPIRE_TIME = time.Minute * 5 // 5 mins

const BASE_URL = "https://bpay.binanceapi.com"
const CREATE_ORDER_PATH = "/binancepay/openapi/v2/order"
const QUERY_ORDER_PATH = "/binancepay/openapi/v2/order/query"

const QUERY_ORDER_STATUS_INITIAL = "INITIAL"
const QUERY_ORDER_STATUS_PAID = "PAID"
const QUERY_ORDER_STATUS_EXPIRED = "EXPIRED"
const QUERY_ORDER_STATUS_PENDING = "PENDING"
const QUERY_ORDER_STATUS_CANCELED = "CANCELED"
const QUERY_ORDER_STATUS_ERROR = "ERROR"
const QUERY_ORDER_STATUS_REFUNDING = "REFUNDING"
const QUERY_ORDER_STATUS_REFUNDED = "REFUNDED"
