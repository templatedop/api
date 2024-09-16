package domain

import (
	"time"

	"github.com/aarondl/opt/null"
	"github.com/jackc/pgtype"
)

type Combinedstruct struct {
	subp  InternationalArticleSubpiece
	asubp []InternationalArticleSubpiece
}
type InternationalArticleSubpiece struct {
	ID                      int       `json:"mailbooking_intl_subpiece_id" db:"mailbooking_intl_subpiece_id"`
	MailBookingIntlID       int64     `json:"mailbooking_intl_id"`
	HSCD                    string    `json:"hs_cd"`
	CTHCD                   string    `json:"cth_cd"`
	HSDescription           string    `json:"hs_description"`
	SPUnitCD                string    `json:"sp_unit_cd,"`
	SPCount                 int       `json:"sp_count"`
	SPWeightTotal           int       `json:"sp_weight_total"`
	SPWeightNett            int       `json:"sp_weight_nett"`
	SPOriginCurrencyCD      string    `json:"sp_origin_currency_cd"`
	SPCommInvoiceNo         string    `json:"sp_comm_invoice_no"`
	SPCommInvoiceDate       time.Time `json:"sp_comm_invoice_date"`
	SPInvCurrencyCD         string    `json:"sp_inv_currency_cd"`
	SPInvCurrencyExchrate   int       `json:"sp_inv_currency_exchrate"`
	SPAsblFOBValue          int       `json:"sp_asbl_fob_value"`
	SPAsblValueINR          int       `json:"sp_asbl_value_inr"`
	SPTaxInvoiceNo          string    `json:"sp_tax_invoice_no"`
	SPTaxInvoiceDate        time.Time `json:"sp_tax_invoice_date"`
	SPInvoiceLSN            int       `json:"sp_invoice_lsn"`
	SPInvoiceValuePU        int       `json:"sp_invoice_value_pu"`
	SPInvoiceValueTotal     int       `json:"sp_invoice_value_total"`
	ECommerceURL            string    `json:"ecommerce_url"`
	ECommercePaytranID      string    `json:"ecommerce_paytranid"`
	ECommerceSKU            string    `json:"ecommerce_sku"`
	CreatedOn               time.Time `json:"createdon"`
	CreatedBy               string    `json:"createdby"`
	CounterNo               int       `json:"counterno"`
	ShiftNo                 int       `json:"shiftno"`
	FacilityIDBKG           string    `json:"facilityid_bkg"`
	IPAddressBKG            string    `json:"ipaddress_bkg"`
	UpdatedOn               time.Time `json:"updatedon"`
	UpdatedBy               string    `json:"updatedby"`
	FacilityIDUPD           string    `json:"facilityid_upd"`
	IPAddressUPD            string    `json:"ipaddress_upd"`
	UserTypeCD              string    `json:"usertype_cd"`
	ChannelTypeCD           string    `json:"channeltype_cd"`
	IGSTRate                float64   `json:"igst_rate"`
	IGSTAmount              float64   `json:"igst_amount"`
	ExportDutyRate          float64   `json:"export_duty_rate"`
	ExportDutyAmount        float64   `json:"export_duty_amount"`
	CessRate                float64   `json:"cess_rate"`
	CessAmount              float64   `json:"cess_amount"`
	CompensationCessRate    float64   `json:"compensation_cess_rate"`
	CompensationCessAmount  float64   `json:"compensation_cess_amount"`
	TaxPaymentModeCD        string    `json:"tax_payment_mode_cd"`
	TaxPaymentChannelRefNo  int       `json:"tax_payment_channel_ref_no"`
	TaxPaymentChannelDate   time.Time `json:"tax_payment_channel_date"`
	TaxPaymentChannelSource string    `json:"tax_payment_channel_source"`
}

type InternationalArticleSubpiecedb struct {
	ID                      int       `db:"mailbooking_intl_subpiece_id" `
	MailBookingIntlID       int64     `db:"mailbooking_intl_id"`
	HSCD                    string    `db:"hs_cd"`
	CTHCD                   string    `db:"cth_cd"`
	HSDescription           string    `db:"hs_description"`
	SPUnitCD                string    `db:"sp_unit_cd,"`
	SPCount                 int       `db:"sp_count"`
	SPWeightTotal           int       `db:"sp_weight_total"`
	SPWeightNett            int       `db:"sp_weight_nett"`
	SPOriginCurrencyCD      string    `db:"sp_origin_currency_cd"`
	SPCommInvoiceNo         string    `db:"sp_comm_invoice_no"`
	SPCommInvoiceDate       time.Time `db:"sp_comm_invoice_date"`
	SPInvCurrencyCD         string    `db:"sp_inv_currency_cd"`
	SPInvCurrencyExchrate   int       `db:"sp_inv_currency_exchrate"`
	SPAsblFOBValue          int       `db:"sp_asbl_fob_value"`
	SPAsblValueINR          int       `db:"sp_asbl_value_inr"`
	SPTaxInvoiceNo          string    `db:"sp_tax_invoice_no"`
	SPTaxInvoiceDate        time.Time `db:"sp_tax_invoice_date"`
	SPInvoiceLSN            int       `db:"sp_invoice_lsn"`
	SPInvoiceValuePU        int       `db:"sp_invoice_value_pu"`
	SPInvoiceValueTotal     int       `db:"sp_invoice_value_total"`
	ECommerceURL            string    `db:"ecommerce_url"`
	ECommercePaytranID      string    `db:"ecommerce_paytranid"`
	ECommerceSKU            string    `db:"ecommerce_sku"`
	CreatedOn               time.Time `db:"createdon"`
	CreatedBy               string    `db:"createdby"`
	CounterNo               int       `db:"counterno"`
	ShiftNo                 int       `db:"shiftno"`
	FacilityIDBKG           string    `db:"facilityid_bkg"`
	IPAddressBKG            string    `db:"ipaddress_bkg"`
	UpdatedOn               time.Time `db:"updatedon"`
	UpdatedBy               string    `db:"updatedby"`
	FacilityIDUPD           string    `db:"facilityid_upd"`
	IPAddressUPD            string    `db:"ipaddress_upd"`
	UserTypeCD              string    `db:"usertype_cd"`
	ChannelTypeCD           string    `db:"channeltype_cd"`
	IGSTRate                float64   `db:"igst_rate"`
	IGSTAmount              float64   `db:"igst_amount"`
	ExportDutyRate          float64   `db:"export_duty_rate"`
	ExportDutyAmount        float64   `db:"export_duty_amount"`
	CessRate                float64   `db:"cess_rate"`
	CessAmount              float64   `db:"cess_amount"`
	CompensationCessRate    float64   `db:"compensation_cess_rate"`
	CompensationCessAmount  float64   `db:"compensation_cess_amount"`
	TaxPaymentModeCD        string    `db:"tax_payment_mode_cd"`
	TaxPaymentChannelRefNo  int       `db:"tax_payment_channel_ref_no"`
	TaxPaymentChannelDate   time.Time `db:"tax_payment_channel_date"`
	TaxPaymentChannelSource string    `db:"tax_payment_channel_source"`
}

type ISubpieces struct {
	IntlSubpieces []InternationalArticleSubpiece `json:"subpieces"`
}

type Phone struct {
	Number string `json:"number"`
	Type   string `json:"type"`
}

type Bag struct {
	BagID     int         `json:"bagid"`
	BagName   pgtype.Text `json:"bagname"`
	BagWeight float64     `json:"bagweight"`
	Articles  []Article   `json:"articles"`
	Phones    []Phone     `json:"phones"`
}

type Bag1 struct {
	BagID     int              `json:"bagid"`
	BagName   null.Val[string] `json:"bagname"`
	BagWeight float64          `json:"bagweight"`
	//Testjson  pgtype.JSON      `json:"testjson"`

	//Test      string           `json:"test"`
}

type Article struct {
	ArticleID int    `json:"articleid"`
	Address   string `json:"address"`
}
type Bags struct {
	Bags []Bag `json:"bags,omitempty"`
}
