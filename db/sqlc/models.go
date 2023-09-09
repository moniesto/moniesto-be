// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type BinancePaymentDateType string

const (
	BinancePaymentDateTypeMONTH BinancePaymentDateType = "MONTH"
)

func (e *BinancePaymentDateType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = BinancePaymentDateType(s)
	case string:
		*e = BinancePaymentDateType(s)
	default:
		return fmt.Errorf("unsupported scan type for BinancePaymentDateType: %T", src)
	}
	return nil
}

type NullBinancePaymentDateType struct {
	BinancePaymentDateType BinancePaymentDateType `json:"binance_payment_date_type"`
	Valid                  bool                   `json:"valid"` // Valid is true if BinancePaymentDateType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullBinancePaymentDateType) Scan(value interface{}) error {
	if value == nil {
		ns.BinancePaymentDateType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.BinancePaymentDateType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullBinancePaymentDateType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.BinancePaymentDateType), nil
}

type BinancePaymentStatus string

const (
	BinancePaymentStatusPending BinancePaymentStatus = "pending"
	BinancePaymentStatusFail    BinancePaymentStatus = "fail"
	BinancePaymentStatusSuccess BinancePaymentStatus = "success"
)

func (e *BinancePaymentStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = BinancePaymentStatus(s)
	case string:
		*e = BinancePaymentStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for BinancePaymentStatus: %T", src)
	}
	return nil
}

type NullBinancePaymentStatus struct {
	BinancePaymentStatus BinancePaymentStatus `json:"binance_payment_status"`
	Valid                bool                 `json:"valid"` // Valid is true if BinancePaymentStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullBinancePaymentStatus) Scan(value interface{}) error {
	if value == nil {
		ns.BinancePaymentStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.BinancePaymentStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullBinancePaymentStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.BinancePaymentStatus), nil
}

type BinancePayoutStatus string

const (
	BinancePayoutStatusPending    BinancePayoutStatus = "pending"
	BinancePayoutStatusFail       BinancePayoutStatus = "fail"
	BinancePayoutStatusSuccess    BinancePayoutStatus = "success"
	BinancePayoutStatusRefund     BinancePayoutStatus = "refund"
	BinancePayoutStatusRefundFail BinancePayoutStatus = "refund_fail"
)

func (e *BinancePayoutStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = BinancePayoutStatus(s)
	case string:
		*e = BinancePayoutStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for BinancePayoutStatus: %T", src)
	}
	return nil
}

type NullBinancePayoutStatus struct {
	BinancePayoutStatus BinancePayoutStatus `json:"binance_payout_status"`
	Valid               bool                `json:"valid"` // Valid is true if BinancePayoutStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullBinancePayoutStatus) Scan(value interface{}) error {
	if value == nil {
		ns.BinancePayoutStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.BinancePayoutStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullBinancePayoutStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.BinancePayoutStatus), nil
}

type EntryPosition string

const (
	EntryPositionLong  EntryPosition = "long"
	EntryPositionShort EntryPosition = "short"
)

func (e *EntryPosition) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = EntryPosition(s)
	case string:
		*e = EntryPosition(s)
	default:
		return fmt.Errorf("unsupported scan type for EntryPosition: %T", src)
	}
	return nil
}

type NullEntryPosition struct {
	EntryPosition EntryPosition `json:"entry_position"`
	Valid         bool          `json:"valid"` // Valid is true if EntryPosition is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullEntryPosition) Scan(value interface{}) error {
	if value == nil {
		ns.EntryPosition, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.EntryPosition.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullEntryPosition) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.EntryPosition), nil
}

type ImageType string

const (
	ImageTypeProfilePhoto    ImageType = "profile_photo"
	ImageTypeBackgroundPhoto ImageType = "background_photo"
)

func (e *ImageType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ImageType(s)
	case string:
		*e = ImageType(s)
	default:
		return fmt.Errorf("unsupported scan type for ImageType: %T", src)
	}
	return nil
}

type NullImageType struct {
	ImageType ImageType `json:"image_type"`
	Valid     bool      `json:"valid"` // Valid is true if ImageType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullImageType) Scan(value interface{}) error {
	if value == nil {
		ns.ImageType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ImageType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullImageType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ImageType), nil
}

type PayoutSource string

const (
	PayoutSourceBINANCE PayoutSource = "BINANCE"
)

func (e *PayoutSource) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = PayoutSource(s)
	case string:
		*e = PayoutSource(s)
	default:
		return fmt.Errorf("unsupported scan type for PayoutSource: %T", src)
	}
	return nil
}

type NullPayoutSource struct {
	PayoutSource PayoutSource `json:"payout_source"`
	Valid        bool         `json:"valid"` // Valid is true if PayoutSource is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullPayoutSource) Scan(value interface{}) error {
	if value == nil {
		ns.PayoutSource, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.PayoutSource.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullPayoutSource) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.PayoutSource), nil
}

type PayoutType string

const (
	PayoutTypeBINANCEID PayoutType = "BINANCE_ID"
)

func (e *PayoutType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = PayoutType(s)
	case string:
		*e = PayoutType(s)
	default:
		return fmt.Errorf("unsupported scan type for PayoutType: %T", src)
	}
	return nil
}

type NullPayoutType struct {
	PayoutType PayoutType `json:"payout_type"`
	Valid      bool       `json:"valid"` // Valid is true if PayoutType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullPayoutType) Scan(value interface{}) error {
	if value == nil {
		ns.PayoutType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.PayoutType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullPayoutType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.PayoutType), nil
}

type PostCryptoMarketType string

const (
	PostCryptoMarketTypeSpot    PostCryptoMarketType = "spot"
	PostCryptoMarketTypeFutures PostCryptoMarketType = "futures"
)

func (e *PostCryptoMarketType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = PostCryptoMarketType(s)
	case string:
		*e = PostCryptoMarketType(s)
	default:
		return fmt.Errorf("unsupported scan type for PostCryptoMarketType: %T", src)
	}
	return nil
}

type NullPostCryptoMarketType struct {
	PostCryptoMarketType PostCryptoMarketType `json:"post_crypto_market_type"`
	Valid                bool                 `json:"valid"` // Valid is true if PostCryptoMarketType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullPostCryptoMarketType) Scan(value interface{}) error {
	if value == nil {
		ns.PostCryptoMarketType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.PostCryptoMarketType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullPostCryptoMarketType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.PostCryptoMarketType), nil
}

type PostCryptoStatus string

const (
	PostCryptoStatusPending PostCryptoStatus = "pending"
	PostCryptoStatusFail    PostCryptoStatus = "fail"
	PostCryptoStatusSuccess PostCryptoStatus = "success"
)

func (e *PostCryptoStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = PostCryptoStatus(s)
	case string:
		*e = PostCryptoStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for PostCryptoStatus: %T", src)
	}
	return nil
}

type NullPostCryptoStatus struct {
	PostCryptoStatus PostCryptoStatus `json:"post_crypto_status"`
	Valid            bool             `json:"valid"` // Valid is true if PostCryptoStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullPostCryptoStatus) Scan(value interface{}) error {
	if value == nil {
		ns.PostCryptoStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.PostCryptoStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullPostCryptoStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.PostCryptoStatus), nil
}

type UserLanguage string

const (
	UserLanguageEn UserLanguage = "en"
	UserLanguageTr UserLanguage = "tr"
)

func (e *UserLanguage) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserLanguage(s)
	case string:
		*e = UserLanguage(s)
	default:
		return fmt.Errorf("unsupported scan type for UserLanguage: %T", src)
	}
	return nil
}

type NullUserLanguage struct {
	UserLanguage UserLanguage `json:"user_language"`
	Valid        bool         `json:"valid"` // Valid is true if UserLanguage is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUserLanguage) Scan(value interface{}) error {
	if value == nil {
		ns.UserLanguage, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UserLanguage.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUserLanguage) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UserLanguage), nil
}

// Stores binance payment transactions info and history
type BinancePaymentTransaction struct {
	ID            string                 `json:"id"`
	QrcodeLink    string                 `json:"qrcode_link"`
	CheckoutLink  string                 `json:"checkout_link"`
	DeepLink      string                 `json:"deep_link"`
	UniversalLink string                 `json:"universal_link"`
	Status        BinancePaymentStatus   `json:"status"`
	UserID        string                 `json:"user_id"`
	MoniestID     string                 `json:"moniest_id"`
	DateType      BinancePaymentDateType `json:"date_type"`
	DateValue     int32                  `json:"date_value"`
	MoniestFee    float64                `json:"moniest_fee"`
	Amount        float64                `json:"amount"`
	WebhookUrl    string                 `json:"webhook_url"`
	PayerID       sql.NullString         `json:"payer_id"`
	CreatedAt     time.Time              `json:"created_at"`
	UpdatedAt     time.Time              `json:"updated_at"`
}

// Stores binance payout info and history
type BinancePayoutHistory struct {
	ID                     string                 `json:"id"`
	TransactionID          string                 `json:"transaction_id"`
	UserID                 string                 `json:"user_id"`
	MoniestID              string                 `json:"moniest_id"`
	PayerID                string                 `json:"payer_id"`
	TotalAmount            float64                `json:"total_amount"`
	Amount                 float64                `json:"amount"`
	DateType               BinancePaymentDateType `json:"date_type"`
	DateValue              int32                  `json:"date_value"`
	DateIndex              int32                  `json:"date_index"`
	PayoutDate             time.Time              `json:"payout_date"`
	PayoutYear             int32                  `json:"payout_year"`
	PayoutMonth            int32                  `json:"payout_month"`
	PayoutDay              int32                  `json:"payout_day"`
	Status                 BinancePayoutStatus    `json:"status"`
	OperationFeePercentage sql.NullFloat64        `json:"operation_fee_percentage"`
	PayoutDoneAt           sql.NullTime           `json:"payout_done_at"`
	PayoutRequestID        sql.NullString         `json:"payout_request_id"`
	FailureMessage         sql.NullString         `json:"failure_message"`
	CreatedAt              time.Time              `json:"created_at"`
	UpdatedAt              time.Time              `json:"updated_at"`
}

// Stores email verification token for verifying account
type EmailVerificationToken struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Token       string    `json:"token"`
	TokenExpiry time.Time `json:"token_expiry"`
	RedirectUrl string    `json:"redirect_url"`
	Deleted     bool      `json:"deleted"`
	CreatedAt   time.Time `json:"created_at"`
}

// Stores feedback from users
type Feedback struct {
	ID        string         `json:"id"`
	UserID    sql.NullString `json:"user_id"`
	Type      sql.NullString `json:"type"`
	Message   string         `json:"message"`
	Solved    bool           `json:"solved"`
	CreatedAt time.Time      `json:"created_at"`
	SolvedAt  sql.NullTime   `json:"solved_at"`
}

// Stores image data
type Image struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	Link          string    `json:"link"`
	ThumbnailLink string    `json:"thumbnail_link"`
	Type          ImageType `json:"type"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// Stores moniest data
type Moniest struct {
	ID          string         `json:"id"`
	UserID      string         `json:"user_id"`
	Bio         sql.NullString `json:"bio"`
	Description sql.NullString `json:"description"`
	Deleted     bool           `json:"deleted"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

// Stores moniest payout info
type MoniestPayoutInfo struct {
	ID        string       `json:"id"`
	MoniestID string       `json:"moniest_id"`
	Source    PayoutSource `json:"source"`
	Type      PayoutType   `json:"type"`
	Value     string       `json:"value"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

// Stores moniest crypto statistics info
type MoniestPostCryptoStatistic struct {
	ID            string          `json:"id"`
	MoniestID     string          `json:"moniest_id"`
	Pnl7days      sql.NullFloat64 `json:"pnl_7days"`
	Roi7days      sql.NullFloat64 `json:"roi_7days"`
	WinRate7days  sql.NullFloat64 `json:"win_rate_7days"`
	Posts7days    []string        `json:"posts_7days"`
	Pnl30days     sql.NullFloat64 `json:"pnl_30days"`
	Roi30days     sql.NullFloat64 `json:"roi_30days"`
	WinRate30days sql.NullFloat64 `json:"win_rate_30days"`
	Posts30days   []string        `json:"posts_30days"`
	PnlTotal      sql.NullFloat64 `json:"pnl_total"`
	RoiTotal      sql.NullFloat64 `json:"roi_total"`
	WinRateTotal  sql.NullFloat64 `json:"win_rate_total"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
}

// Stores subscription data of a moniest
type MoniestSubscriptionInfo struct {
	ID        string         `json:"id"`
	MoniestID string         `json:"moniest_id"`
	Fee       float64        `json:"fee"`
	Message   sql.NullString `json:"message"`
	Deleted   bool           `json:"deleted"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

// Stores reset token for forget password operations
type PasswordResetToken struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Token       string    `json:"token"`
	TokenExpiry time.Time `json:"token_expiry"`
	Deleted     bool      `json:"deleted"`
	CreatedAt   time.Time `json:"created_at"`
}

// Stores crypto posts data
type PostCrypto struct {
	ID             string               `json:"id"`
	MoniestID      string               `json:"moniest_id"`
	MarketType     PostCryptoMarketType `json:"market_type"`
	Currency       string               `json:"currency"`
	StartPrice     float64              `json:"start_price"`
	Duration       time.Time            `json:"duration"`
	TakeProfit     float64              `json:"take_profit"`
	Stop           float64              `json:"stop"`
	Target1        sql.NullFloat64      `json:"target1"`
	Target2        sql.NullFloat64      `json:"target2"`
	Target3        sql.NullFloat64      `json:"target3"`
	Direction      EntryPosition        `json:"direction"`
	Leverage       int32                `json:"leverage"`
	Finished       bool                 `json:"finished"`
	Status         PostCryptoStatus     `json:"status"`
	Pnl            float64              `json:"pnl"`
	Roi            float64              `json:"roi"`
	LastOperatedAt time.Time            `json:"last_operated_at"`
	Deleted        bool                 `json:"deleted"`
	CreatedAt      time.Time            `json:"created_at"`
	UpdatedAt      time.Time            `json:"updated_at"`
}

// Stores crypto post description data
type PostCryptoDescription struct {
	ID          string    `json:"id"`
	PostID      string    `json:"post_id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Stores user data
type User struct {
	ID            string         `json:"id"`
	Fullname      string         `json:"fullname"`
	Username      string         `json:"username"`
	Email         string         `json:"email"`
	EmailVerified bool           `json:"email_verified"`
	Password      string         `json:"password"`
	Location      sql.NullString `json:"location"`
	LoginCount    int32          `json:"login_count"`
	Language      UserLanguage   `json:"language"`
	Deleted       bool           `json:"deleted"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	LastLogin     time.Time      `json:"last_login"`
}

// Stores user subscription info
type UserSubscription struct {
	ID                    string         `json:"id"`
	UserID                string         `json:"user_id"`
	MoniestID             string         `json:"moniest_id"`
	Active                bool           `json:"active"`
	LatestTransactionID   sql.NullString `json:"latest_transaction_id"`
	SubscriptionStartDate time.Time      `json:"subscription_start_date"`
	SubscriptionEndDate   time.Time      `json:"subscription_end_date"`
	CreatedAt             time.Time      `json:"created_at"`
	UpdatedAt             time.Time      `json:"updated_at"`
}

// Stores user subscriptions history
type UserSubscriptionHistory struct {
	ID                    string         `json:"id"`
	UserID                string         `json:"user_id"`
	MoniestID             string         `json:"moniest_id"`
	TransactionID         sql.NullString `json:"transaction_id"`
	SubscriptionStartDate time.Time      `json:"subscription_start_date"`
	SubscriptionEndDate   time.Time      `json:"subscription_end_date"`
	CreatedAt             time.Time      `json:"created_at"`
}
