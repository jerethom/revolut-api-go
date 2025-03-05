package revolut_merchant

import (
	"encoding/json"
	"fmt"
	"time"
)

type OrderCustomer struct {
	// Permanent ID of a customer used to retrieve, update, delete a customer. This ID can also be used to link customer to an order.
	Id string `json:"id,omitempty"`
	//Possible values: >= 2 characters
	//
	//The customer's full name.
	FullName string `json:"full_name,omitempty"`
	// The customer's phone number.
	Phone string `json:"phone,omitempty"`
	// The customer's email address.
	Email string `json:"email,omitempty"`
}

type Address struct {
	//Possible values: <= 100 characters
	//
	//Street line 1 information.
	StreetLine1 string `json:"street_line_1,omitempty"`
	//Possible values: <= 100 characters
	//
	//Street line 2 information.
	StreetLine2 string `json:"street_line_2,omitempty"`
	//Possible values: <= 100 characters
	//
	//The region associated with the address.
	Region string `json:"region,omitempty"`
	//Possible values: <= 100 characters
	//
	//The city associated with the address.
	City string `json:"city,omitempty"`
	//Possible values: <= 2 characters
	//
	//The 2-letter country code of the country associated with the address.
	CountryCode string `json:"country_code"`
	//Possible values: <= 100 characters
	//
	//The postcode associated with the address.
	Postcode string `json:"postcode"`
}

type Quantity struct {
	// The number of units of the line item.
	Value int `json:"value"`
	//Possible values: <= 100 characters
	//
	//The measurement unit for the quantity, such as cm, or kg.
	Unit string `json:"unit,omitempty"`
}
type Discount struct {
	//Possible values: <= 100 characters
	//
	//The specific name or label of the discount applied to the line item.
	Name string `json:"name"`
	// The monetary value of the discount.
	Amount int `json:"amount"`
}

type Taxe struct {
	// The specific name or designation of the tax applied to the line item.
	Name string `json:"name"`
	// The monetary value of the tax.
	Amount int `json:"amount"`
}

type LineItem struct {
	//Possible values: <= 250 characters
	//
	//Name of the line item.
	Name string `json:"name"`
	//Possible values: [physical, service]
	//
	//Type of the line item.
	Type string `json:"type"`
	// Object representing the quantity details of a line item, including the amount and its associated unit of measurement.
	Quantity Quantity `json:"quantity"`
	// The unit price of the line item.
	UnitPriceAmount int `json:"unit_price_amount"`
	// The total amount to be paid for the line item, including taxes and discounts.
	TotalAmount int `json:"total_amount"`
	//Possible values: <= 250 characters
	//
	//Unique identifier of line item in the merchant's system.
	ExternalId string `json:"external_id,omitempty"`
	//Possible values: <= 50
	//
	//A list of discounts applied to the line item. Each discount should be subtracted from the total amount payable for the item.
	Discounts []Discount `json:"discounts,omitempty"`
	//Possible values: <= 50
	//
	//A list of taxes applied to the line item. Each tax should be added to the total amount payable for the item.
	Taxes []Taxe `json:"taxes,omitempty"`
	//Possible values: <= 50
	//
	//A list of URLs pointing to images related to the line item. These images can provide visual details or representations of the item.
	ImageUrls []string `json:"image_urls,omitempty"`
	//Possible values: <= 1024 characters
	//
	//Description of the line item.
	Description string `json:"description,omitempty"`
	//Possible values: <= 2000 characters
	//
	//An HTTP/HTTPS URL that links to more information about the line item, such as a product page or details.
	Url string `json:"url,omitempty"`
}

type MerchantOrderData struct {
	//Possible values: <= 2000 characters
	//
	//The URL of the order stored in the merchant's order management system.
	//
	//This URL will be included in the order confirmation email for payments made via Revolut Pay. If specified, this URL will override the default link to the merchant's Business website configured in the Revolut Business account.
	Url string `json:"url,omitempty"`
	//Merchant order ID for external reference.
	//
	//Use this field to set the ID that your own system can use to easily track orders.
	Reference string `json:"reference,omitempty"`
}

type UpcomingPaymentData struct {
	// The date and time in ISO 8601 format when the upcoming payment is scheduled to be executed.
	Date time.Time `json:"date"`
	// The unique ID of the customer's payment method used to complete the scheduled payment.
	PaymentMethodId string `json:"payment_method_id"`
}

type Contact struct {
	//Possible values: <= 250 characters
	//
	//Full name of the contact person.
	Name string `json:"name,omitempty"`
	// Email address of the contact person.
	Email string `json:"email"`
	// Phone number of the contact person.
	Phone string `json:"phone"`
}

type Shipment struct {
	//Possible values: <= 250 characters
	//
	//Name of the company handling the shipment.
	ShippingCompanyName string `json:"shipping_company_name"`
	//Possible values: <= 500 characters
	//
	//Unique tracking number for the shipment.
	TrackingNumber string `json:"tracking_number"`
	// Estimated delivery date and time for the shipment. The time should also include the customer's timezone.
	EstimatedDeliveryDate time.Time `json:"estimated_delivery_date,omitempty"`
	//Possible values: <= 2000 characters
	//
	//An HTTP/HTTPS URL where the shipment can be tracked.
	TrackingUrl string `json:"tracking_url,omitempty"`
}

type Shipping struct {
	// Details of a physical address.
	Address Address `json:"address,omitempty"`
	// Contact details for someone responsible for the shipment.
	Contact Contact `json:"contact,omitempty"`
	//Possible values: <= 50
	//
	//List of individual shipment details.
	Shipments []Shipment `json:"shipments,omitempty"`
}

type Passenger struct {
	// Passenger's first name.
	FirstName string `json:"first_name"`
	// Passenger's last name.
	LastName string `json:"last_name"`
}

type JourneyLeg struct {
	//Sequence of the journey legs. Increment by 1 for each flight included in the ticket.
	//
	//For example: For a FRA > LHR > DUB > LHR > FRA journey the sequence will be assigned as:
	//
	//FRA > LHR: 1
	//LHR > DUB: 2
	//DUB > LHR: 3
	//LHR > FRA: 4
	Sequence int `json:"sequence"`
	// The IATA 3-letter airport code for the departure airport.
	DepartureAirportCode string `json:"departure_airport_code"`
	// The IATA 3-letter airport code for the arrival airport.
	ArrivalAirportCode string `json:"arrival_airport_code"`
	// The flight identifier, without airline code.
	FlightNumber string `json:"flight_number,omitempty"`
	// The fare base code for the given journey leg.
	FareBaseCode string `json:"fare_base_code,omitempty"`
	// The UTC date and time of the flight departure for the given journey leg.
	TravelDate time.Time `json:"travel_date"`
	// The name of the airline associated with the journey leg.
	AirlineName string `json:"airlinename"`
	// The IATA 2-letter accounting code identifying the airline associated with the journey leg.
	AirlineCode string `json:"airline_code"`
}

type Transaction struct {
	//Possible values: <= 255 characters
	//
	//The public transaction hash of the crypto transaction.
	Id string `json:"id"`
	//Possible values: [pending, failed, cancelled, completed]
	//
	//The status of the crypto transaction.
	Status string `json:"status"`
	//Possible values: <= 255 characters
	//
	//The wallet ID of the recipient.
	RecipientWalletId string `json:"recipient_wallet_id,omitempty"`
	//Possible values: <= 255 characters
	//
	//The user ID of the recipient.
	RecipientUserId string `json:"recipient_user_id,omitempty"`
}

type Airline struct {
	//Possible values: [airline, crypto]
	//
	//Type of the industry-specific data object, determining what further data is expected for a particular order.
	//
	//Available types:
	//
	//airline
	//crypto
	Type string `json:"type"`
	// Unique ID of the travel reservation associated with the order.
	BookingId string `json:"booking_id"`
	// The UTC date and time of the final journey leg.
	FulfillmentDate time.Time `json:"fulfillment_date,omitempty"`
	//Possible values: [flexible, fixed]
	//
	//The type of the ticket.
	TicketType string `json:"ticket_type,omitempty"`
	// The code of the Computer Reservation System (CRS) used to make the booking and purchase the ticket.
	CrsCode string `json:"crs_code,omitempty"`
	//Possible values: [new, modification]
	//
	//Parameter indicating whether this order is related to a new ticket reservation or a modification of an existing one.
	TicketChangeIndicator string `json:"ticket_change_indicator,omitempty"`
	//Possible values: [refundable, non_refundable, partially_refundable]
	//
	//Parameter indicating whether the ticket is refundable, partially refundable, or not refundable.
	Refundability string `json:"refundability,omitempty"`
	// Array containing information of passengers associated with the booking.
	Passengers []Passenger `json:"passengers,omitempty"`
	// Array containing information of journey legs associated with the booking.
	JourneyLegs []JourneyLeg `json:"journey_legs,omitempty"`
}

type Crypto struct {
	//Possible values: [airline, crypto]
	//
	//Type of the industry-specific data object, determining what further data is expected for a particular order.
	//
	//Available types:
	//
	//airline
	//crypto
	Type string `json:"type"`
	//Possible values: >= 1
	//
	//Array of crypto transaction data associated with the order.
	Transactions []Transaction `json:"transactions"`
}

type IndustryData struct {
	// Object containing industry-specific information associated with the order.
	Airline *Airline
	// Object containing industry-specific information associated with the order.
	Crypto *Crypto
}

func (i *IndustryData) UnmarshalJSON(bytes []byte) error {
	var temp map[string]interface{}
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return err
	}

	if typeStr, ok := temp["type"].(string); ok {
		switch typeStr {
		case "airline":
			i.Airline = new(Airline)
			return json.Unmarshal(bytes, i.Airline)
		case "crypto":
			i.Crypto = new(Crypto)
			return json.Unmarshal(bytes, i.Crypto)
		default:
			return fmt.Errorf("unknown type %s", typeStr)
		}
	}

	return fmt.Errorf("missing or invalid 'type' field")
}

func (i IndustryData) MarshalJSON() ([]byte, error) {
	if i.Airline != nil {
		return json.Marshal(i.Airline)
	}
	if i.Crypto != nil {
		return json.Marshal(i.Crypto)
	}
	return json.Marshal(nil)
}

type CreateAnOrderPayload struct {
	// The total amount of the order in minor currency units. For example, 7034 represents €70.34.
	Amount int `json:"amount"`
	// ISO 4217 currency code in upper case.
	Currency string `json:"currency"`
	//ISO 4217 currency code in upper case.
	//
	//If settlement_currency is different from the value of currency, the money will be exchanged when the amount is settled to your merchant account. In case of a refund or chargeback, the money will be exchanged to the order's initial currency.
	//
	//If settlement_currency is not specified, this value is taken from currency.
	SettlementCurrency string `json:"settlement_currency,omitempty"`
	// The description of the order.
	Description string `json:"description,omitempty"`
	//Object containing information about a customer.
	//
	//If you have it, we strongly advise providing at least either id, phone, or email.
	//
	//Using the Customers operations, you can manage customer instances.
	Customer *OrderCustomer `json:"customer,omitempty"`
	//Possible values: [automatic, forced]
	//
	//Default value: automatic
	//
	//The enforce challenge mode. automatic is used by default.
	EnforceChallenge string `json:"enforce_challenge,omitempty"`
	//An array of line items included in the order.
	//
	//Each line item represents an individual product or service, along with its quantity, price, taxes, and discounts.
	LineItems []LineItem `json:"line_items,omitempty"`
	// Details about the shipping related to the order, including address, contact information, and individual shipments.
	Shipping *Shipping `json:"shipping,omitempty"`
	//Possible values: [automatic, manual]
	//
	//Default value: automatic
	//
	//The capture mode of the order. automatic is used by default.
	CaptureMode string `json:"capture_mode,omitempty"`
	//Automatic cancellation period for uncaptured orders, specified in ISO 8601 format.
	//
	//Orders in authorised state will be automatically cancelled if they stay uncaptured for longer than the period specified. Maximum: 7 days = P7D.
	CancelAuthorizedAfter string `json:"cancel_authorized_after,omitempty"`
	// Unique ID representing the location where merchants sells products.
	LocationId string `json:"location_id,omitempty"`
	//Possible values: <= 50
	//
	//Additional information to track your orders in your system, by providing custom metadata using "<key>" : "<value>" pairs.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Object containing industry-specific information associated with the order.
	IndustryData *IndustryData `json:"industry_data,omitempty"`
	// Object containing industry-specific information associated with the order.
	MerchantOrderData *MerchantOrderData `json:"merchant_order_data,omitempty"`
	// Object containing information about upcoming payments associated with the order.
	UpcomingPaymentData *UpcomingPaymentData `json:"upcoming_payment_data,omitempty"`
	// The URL your customer will be redirected to after completing a payment on the hosted checkout page (checkout_url parameter's value of the order).
	RedirectUrl string `json:"redirect_url,omitempty"`
}

type ChecksThreeDs struct {
	// The Electronic Commerce Indicator (ECI) value corresponds to the authentication result and indicates the level of security used when the payment information was provided.
	Eci string `json:"eci"`
	//Possible values: [verified, failed, challenge]
	//
	//The result of 3D Secure check.
	State string `json:"state"`
	// The 3D Secure version.
	Version int `json:"version"`
}

type Checks struct {
	// The details of the 3D Secure check. Only for orders with successful payments.
	ThreeDs *ChecksThreeDs `json:"three_ds"`
	//Possible values: [match, not_match, incorrect, not_processed]
	//
	//The result of CVV verification.
	CvvVerification string `json:"cvv_verification"`
	//Possible values: [match, not_match, n_a, invalid]
	//
	//The result of address verification.
	Address string `json:"address"`
	//Possible values: [match, not_match, n_a, invalid]
	//
	//The result of postcode verification.
	Postcode string `json:"postcode"`
	//Possible values: [match, not_match, n_a, invalid]
	//
	//The result of cardholder verification.
	CardHolder string `json:"card_holder"`
}

type ThreeDs struct {
	//Possible values: [three_ds, three_ds_fingerprint]
	//
	//Type of the authentication challenge the payment triggers.
	Type string `json:"type"`
	// The URL of the authentication challenge.
	AcsUrl string `json:"acs_url"`
}

type ThreeDsFingerPrint struct {
	//Possible values: [three_ds, three_ds_fingerprint]
	//
	//Type of the authentication challenge the payment triggers.
	Type string `json:"type"`
	// The URL of the fingerprint. (Used only internally by Revolut.)
	FingerprintUrl string `json:"fingerprint_url"`
	// Data about the fingerprint used for authentication. (Used only internally by Revolut.)
	FingerPrintData map[string]any `json:"fingerprint_data"`
}

type Card struct {
	// ID of the saved payment method.
	Id string `json:"id"`
	//Possible values: [apple_pay, apple_tap_to_pay, card, cash, google_pay, revolut_pay_card, revolut_pay_account]
	//
	//The type of payment method used to pay for the order.
	Type string `json:"type"`
	//Possible values: [visa, mastercard, american_express]
	//
	//The type of the card.
	CardBrand string `json:"card_brand"`
	//Possible values: [credit, debit, prepaid]
	//
	//The type of card funding.
	Funding string `json:"funding"`
	// The 2-letter country code of the country where the card was issued.
	CardCountryCode string `json:"card_country_code"`
	//Possible values: >= 6 characters and <= 6 characters
	//
	//The BIN of the card.
	CardBin string `json:"card_bin"`
	//Possible values: >= 4 characters and <= 4 characters
	//
	//The last four digits of the card.
	CardLastFour string `json:"card_last_four"`
	// The expiry date of the card in the format of MM/YY.
	CardExpiry string `json:"card_expiry"`
	// The name of the cardholder.
	CardHolderName string `json:"card_holder_name"`
	// The details of the check for card payment. Only for orders with successful payments.
	Checks *Checks `json:"checks"`
}

type PayAccount struct {
	// ID of the saved payment method.
	Id string `json:"id"`
	//Possible values: [apple_pay, apple_tap_to_pay, card, cash, google_pay, revolut_pay_card, revolut_pay_account]
	//
	//The type of payment method used to pay for the order.
	Type string `json:"type"`
}

type Order struct {
	// Permanent order ID used to retrieve, capture, cancel, or refund an order after authorization.
	Id string `json:"id"`
	//Temporary ID for the order, which expires when the payment is authorised.
	//
	//The order token is used to initialise the Revolut Checkout widget, and to be returned by the createOrder callback on the Revolut Pay widget and Apple Pay and Google Pay widget.
	Token string `json:"token"`
	//Possible values: [payment, payment_request, refund, chargeback, chargeback_reversal, credit_reimbursement]
	//
	//The type of the order.
	Type string `json:"type"`
	//Possible values: [pending, processing, authorised, completed, cancelled, failed]
	//
	//The state of the order.
	State string `json:"state"`
	// The date and time the order was created.
	CreatedAt time.Time `json:"created_at"`
	// The date and time the order was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// The description of the order.
	Description string `json:"description"`
	//Possible values: [automatic, manual]
	//
	//Default value: automatic
	//
	//The capture mode of the order. automatic is used by default.
	CaptureMode string `json:"capture_mode"`
	//Automatic cancellation period for uncaptured orders, specified in ISO 8601 format.
	//
	//Orders in authorised state will be automatically cancelled if they stay uncaptured for longer than the period specified. Maximum: 7 days = P7D.
	CancelAuthorizedAfter string `json:"cancel_authorized_after"`
	// The total amount of the order in minor currency units. For example, 7034 represents €70.34.
	Amount int `json:"amount"`
	//The amount not yet paid for a given order (in minor currency units). For example, 7034 represents €70.34.
	//
	//The value in this field may differ from amount if there are partial payments associated with the order.
	OutstandingAmount int `json:"outstanding_amount"`
	//The amount that was refunded from the order (in minor currency units). For example, 7034 represents €70.34.
	//
	//This applies to orders that have been refunded (i.e., orders of type payment that have a related refund order).
	RefundedAmount int `json:"refunded_amount"`
	// ISO 4217 currency code in upper case.
	Currency string `json:"currency"`
	//ISO 4217 currency code in upper case.
	//
	//If settlement_currency is different from the value of currency, the money will be exchanged when the amount is settled to your merchant account. In case of a refund or chargeback, the money will be exchanged to the order's initial currency.
	//
	//If settlement_currency is not specified, this value is taken from currency.
	SettlementAmount int `json:"settlement_amount"`
	//Object containing information about a customer.
	//
	//If you have it, we strongly advise providing at least either id, phone, or email.
	//
	//Using the Customers operations, you can manage customer instances.
	Customer *OrderCustomer `json:"customer"`
	// The details of all the payments that have been made towards this order (successful or unsuccessful).
	Payments []Payment `json:"payments"`
	// Unique ID representing the location where merchants sells products.
	LocationId string `json:"location_id"`
	//Possible values: <= 50
	//
	//Additional information to track your orders in your system, by providing custom metadata using "<key>" : "<value>" pairs.
	Metadata map[string]string `json:"metadata"`
	// Object containing industry-specific information associated with the order.
	IndustryData IndustryData `json:"industry_data"`
	// Object for providing additional information stored in the merchant's order management system.
	MerchantOrderData *MerchantOrderData `json:"merchant_order_data"`
	// Object containing information about upcoming payments associated with the order.
	UpcomingPaymentData *UpcomingPaymentData `json:"upcoming_payment_data"`
	// Link to a checkout page hosted by Revolut.
	CheckoutUrl string `json:"checkout_url"`
	// The URL your customer will be redirected to after completing a payment on the hosted checkout page (checkout_url parameter's value of the order).
	RedirectUrl string `json:"redirect_url"`
	// Details about the shipping related to the order, including address, contact information, and individual shipments.
	Shipping Shipping `json:"shipping"`
	//Possible values: [automatic, forced]
	//
	//Default value: automatic
	//
	//The enforce challenge mode. automatic is used by default.
	EnforceChallenge string `json:"enforce_challenge"`
	//Possible values: <= 250
	//
	//An array of line items included in the order. Each line item represents an individual product or service, along with its quantity, price, taxes, and discounts.
	LineItems []LineItem `json:"line_items"`
}
type CreateAnOrderResponse struct {
	Order
}

type RetrieveAnOrderResponse struct {
	Order
}

type UpdateAnOrderPayload struct {
	// The total amount of the order in minor currency units. For example, 7034 represents €70.34.
	Amount int `json:"amount,omitempty"`
	// ISO 4217 currency code in upper case.
	Currency string `json:"currency,omitempty"`
	//ISO 4217 currency code in upper case.
	//
	//If settlement_currency is different from the value of currency, the money will be exchanged when the amount is settled to your merchant account. In case of a refund or chargeback, the money will be exchanged to the order's initial currency.
	//
	//If settlement_currency is not specified, this value is taken from currency.
	SettlementCurrency string `json:"settlement_currency,omitempty"`
	// The description of the order.
	Description string `json:"description,omitempty"`
	//Object containing information about a customer.
	//
	//If you have it, we strongly advise providing at least either id, phone, or email.
	//
	//Using the Customers operations, you can manage customer instances.
	Customer *OrderCustomer `json:"customer,omitempty"`
	// Details about the shipping related to the order, including address, contact information, and individual shipments.
	Shipping *Shipping `json:"shipping,omitempty"`
	//Possible values: [automatic, forced]
	//
	//Default value: automatic
	//
	//The enforce challenge mode. automatic is used by default.
	EnforceChallenge string `json:"enforce_challenge,omitempty"`
	//Possible values: <= 250
	//
	//An array of line items included in the order. Each line item represents an individual product or service, along with its quantity, price, taxes, and discounts.
	LineItems []LineItem `json:"line_items,omitempty"`
	//Possible values: [automatic, manual]
	//
	//Default value: automatic
	//
	//The capture mode of the order. automatic is used by default.
	CaptureMode string `json:"capture_mode,omitempty"`
	//Automatic cancellation period for uncaptured orders, specified in ISO 8601 format.
	//
	//Orders in authorised state will be automatically cancelled if they stay uncaptured for longer than the period specified. Maximum: 7 days = P7D.
	CancelAuthorizedAfter string `json:"cancel_authorized_after,omitempty"`
	//Possible values: <= 50
	//
	//Additional information to track your orders in your system, by providing custom metadata using "<key>" : "<value>" pairs.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Object containing industry-specific information associated with the order.
	IndustryData *IndustryData `json:"industry_data,omitempty"`
	// Object for providing additional information stored in the merchant's order management system.
	MerchantOrderData *MerchantOrderData `json:"merchant_order_data,omitempty"`
	// Object containing information about upcoming payments associated with the order.
	UpcomingPaymentData *UpcomingPaymentData `json:"upcoming_payment_data,omitempty"`
	// The URL your customer will be redirected to after completing a payment on the hosted checkout page (checkout_url parameter's value of the order).
	RedirectUrl string `json:"redirect_url,omitempty"`
}

type UpdateAnOrderResponse struct {
	Order
}

type CaptureAnOrderPayload struct {
	Amount int `json:"amount,omitempty"`
}

type CaptureAnOrderResponse struct {
	Order
}

type RetrieveAnOrderListResponse []Order

type CancelAnOrderResponse struct {
	Order
}

type RefundAnOrderPayload struct {
	//The amount of the refund (minor currency unit). For example, enter 7034 for €70.34 in the field.
	//
	//This amount can't exceed the remaining amount of the original order. To get the refundable amount, subtract the value of the refunded_amount from the value of the order_amount in the original order. See Retrieve an order.
	Amount int `json:"amount"`
	// The description of the refund.
	Description string `json:"description,omitempty"`
	//Merchant order ID for external reference.
	//
	//Use this field to set the ID that your own system can use to easily track orders.
	MerchantOrderExtRef string `json:"merchant_order_ext_ref,omitempty"`
	//Possible values: <= 50
	//
	//Additional information to track your orders in your system, by providing custom metadata using "<key>" : "<value>" pairs.
	Metadata map[string]string `json:"metadata,omitempty"`
}

type Amount struct {
	// The amount of the settled payment (minor currency unit). For example, enter 7034 for €70.34 in the field.
	Value string `json:"value"`
	// The currency of the settled payment. ISO 4217 currency code in upper case.
	Currency string `json:"currency"`
}

type RelatedOrder struct {
	// The ID of the related order. You can use this ID to get more information about the related order using the Retrieve an order operation.
	Id string `json:"id"`
	//Possible values: [PAYMENT, REFUND, CHARGEBACK]
	//
	//The type of the related order.
	Type string `json:"type"`
	//Possible values: [PENDING, PROCESSING, AUTHORISED, COMPLETED, CANCELLED, FAILED]
	//
	//The state of the related order.
	State string `json:"state"`
	// The amount and currency of the related order.
	Amount *Amount `json:"amount"`
}

type RefundAnOrderResponse struct {
	// Permanent order ID used to retrieve, capture, cancel, or refund an order after authorization.
	Id string `json:"id"`
	//Temporary ID for the order.
	//
	//It expires when the payment is authorized.
	PublicId string `json:"public_id"`
	//Possible values: [PAYMENT, REFUND, CHARGEBACK]
	//
	//The type of the order.
	Type string `json:"type"`
	//Possible values: [PENDING, PROCESSING, AUTHORISED, COMPLETED, CANCELLED, FAILED]
	//
	//The state of the order.
	State string `json:"state"`
	// The date and time the order was created.
	CreatedAt time.Time `json:"created_at"`
	// The date and time the order was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// The date and time the order was completed.
	CompletedAt time.Time `json:"completed_at"`
	// The description of the order.
	Description string `json:"description"`
	//Possible values: [AUTOMATIC, MANUAL]
	//
	//The capture mode of the order. AUTOMATIC is used by default.
	//
	//AUTOMATIC: The order is captured automatically after payment authorisation.
	//MANUAL: The order is not captured automatically. You must manually capture the order later.
	//For more information, see Capture an order.
	CaptureMode string `json:"capture_mode"`
	//ISO 4217 currency code in upper case.
	//
	//If settlement_currency is different from the value of currency, the money will be exchanged when the amount is settled to your merchant account. In case of a refund or chargeback, the money will be exchanged to the order's initial currency.
	//
	//If settlement_currency is not specified, this value is taken from currency.
	SettlementCurrency string `json:"settlement_currency"`
	//Merchant order ID for external reference.
	//
	//Use this field to set the ID that your own system can use to easily track orders.
	MerchantOrderExtRef string `json:"merchant_order_ext_ref"`
	// The ID of the customer associated with this order.
	CustomerId string `json:"customer_id"`
	// The email of the customer.
	Email string `json:"email"`
	// The phone number of the customer.
	Phone string `json:"phone"`
	// The amount and currency of the order.
	OrderAmount *Amount `json:"order_amount"`
	// The amount and currency outstanding to be paid for this order.
	OrderOutstandingAmount *Amount `json:"order_outstanding_amount"`
	// The amount and currency of the refunded order.
	RefundedAmount *Amount `json:"refunded_amount"`
	//
	ShippingAddress *Address `json:"shipping_address"`
	// The details of all the payment attempts that have been made towards this order (successful or unsuccessful).
	Payments []Payment `json:"payments"`
	// The details of related orders. You can use the ID of the related order to Retrieve the order information.
	Related []RelatedOrder `json:"related_orders"`
	//Possible values: <= 50
	//
	//Additional information to track your orders in your system, by providing custom metadata using "<key>" : "<value>" pairs.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Link to a checkout page hosted by Revolut.
	CheckoutUrl string `json:"checkout_url"`
	//Possible values: <= 2000 characters
	//
	//The URI of the order stored in the merchant's order management system.
	//
	//This URI will be included in the order confirmation email for payments made via Revolut Pay. If specified, this URI will override the default link to the merchant's Business website configured in the Revolut Business account.
	MerchantOrderUri string `json:"merchant_order_uri"`
}

type Environment struct {
	//Possible values: [browser]
	//
	//Type of environment where the payment was made.
	Type string `json:"type"`
	//Defines the offset to UTC in minutes.
	TimeZoneUtcOffset int `json:"time_zone_utc_offset"`
	//The browser's available colour depth.
	ColorDepth int `json:"color_depth"`
	//The browser's screen width in pixels.
	ScreenWidth int `json:"screen_width"`
	//The browser's screen height in pixels.
	ScreenHeight int `json:"screen_height"`
	//Indicates if the browser has Java enabled.
	JavaEnabled bool `json:"java_enabled"`
	//Defines the width of the pop-up window where the authentication challenge appears.
	ChallengeWindowWidth int `json:"challenge_window_width"`
	//The URL of the page where the payment was initiated.
	BrowserUrl string `json:"browser_url"`
}

type SavedPaymentMethod struct {
	// Saved payment method ID.
	Id string `json:"id"`
	//Possible values: [card, revolut_pay]
	//
	//Type of saved payment method.
	Type string `json:"type"`
	//Possible values: [customer, merchant]
	//
	//Indicates who is allowed to initiate the payment.
	Initiator string `json:"initiator"`
	// Environment object, indicating in which environment the payment was made.
	Environment Environment `json:"environment"`
}

type PayForAnOrderPayload struct {
	// Object containing information about the saved payment method used to pay for the order.
	SavedPaymentMethod SavedPaymentMethod `json:"saved_payment_method"`
}

type PayForAnOrderResponse struct {
	// The ID of the payment.
	Id string `json:"id"`
	// Permanent order ID used to retrieve, capture, cancel, or refund an order after authorization.
	OrderId string `json:"order_id"`
	// The payment method used to pay for the order.
	// todo PaymentMethod is not the right type for this response see https://developer.revolut.com/docs/merchant/pay-order#response
	PaymentMethod *PaymentMethod `json:"saved_payment_method"`
	//Possible values: [pending, authentication_challenge, authentication_verified, authorisation_started, authorisation_passed, authorised, capture_started, captured, refund_validated, refund_started, cancellation_started, declining, completing, cancelling, failing, completed, declined, soft_declined, cancelled, failed]
	//
	//The status of the payment.
	State string `json:"state"`
	//Details about the authentication challenge that should be performed to complete the authentication process. For more information about Revolut's 3DS solution, see: 3D Secure overview.
	//
	//Only returned if the payment's state is authentication_challenge.
	AuthenticationChallenge *AuthenticationChallenge `json:"authentication_challenge"`
}

type RetrievePaymentListOfAnOrderResponse struct {
	//The ID of the payment.
	Id string `json:"id"`
	//Permanent order ID used to retrieve, capture, cancel, or refund an order after authorization.
	OrderId string `json:"order_id"`
	//The payment method used to pay for the order.
	PaymentMethod *PaymentMethod `json:"saved_payment_method"`
	//Possible values: [pending, authentication_challenge, authentication_verified, authorisation_started, authorisation_passed, authorised, capture_started, captured, refund_validated, refund_started, cancellation_started, declining, completing, cancelling, failing, completed, declined, soft_declined, cancelled, failed]
	//
	//The status of the payment.
	State string `json:"state"`
	//Details about the authentication challenge that should be performed to complete the authentication process. For more information about Revolut's 3DS solution, see: 3D Secure overview.
	//
	//Only returned if the payment's state is authentication_challenge.
	AuthenticationChallenge *AuthenticationChallenge `json:"authentication_challenge"`
}
