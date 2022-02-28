package osipparser2_go

/**

 */

/**
SIP 消息结构，用于请求或应答
*/
type Osip_message_t struct {
	Sip_version                string      /**< SIP version (SIP request only) */
	Req_uri                    *Osip_uri_t /**< Request-Uri (SIP request only) */
	Sip_method                 string      /**< METHOD (SIP request only) */
	Status_code                int
	Reason_phrase              string
	Accepts                    Osip_list_t
	Accept_encodings           Osip_list_t
	Accept_languages           Osip_list_t
	Alert_infos                Osip_list_t
	Allows                     Osip_list_t
	Authentication_infos       Osip_list_t
	Authorizations             Osip_list_t
	Call_id                    *Osip_call_id_t
	Call_infos                 Osip_list_t
	Contacts                   Osip_list_t
	Content_encodings          Osip_list_t
	Content_length             *Osip_content_length_t
	Content_type               *Osip_content_type_t
	Cseq                       *Osip_cseq_t
	Error_infos                Osip_list_t
	From                       *Osip_from_t
	Mime_version               *Osip_mime_version_t
	Proxy_authenticates        Osip_list_t
	Proxy_authentication_infos Osip_list_t
	Proxy_authorizations       Osip_list_t
	Record_routes              Osip_list_t
	Routes                     Osip_list_t
	To                         *Osip_to_t
	Vias                       Osip_list_t
	Www_authenticates          Osip_list_t
	Headers                    Osip_list_t
	Bodies                     Osip_list_t
	Message_property           int32
	Message                    string
	message_length             uint32
	Application_data           []byte
}

const (
	SIP_MESSAGE_MAX_LENGTH uint32 = 8000
	BODY_MESSAGE_MAX_SIZE  uint32 = 4000
)

func Osip_message_init(sip **Osip_message_t) int {
	return 0
}

func Osip_message_free(sip *Osip_message_t) {
	// Do nothing
}

func Osip_message_parse(sip *Osip_message_t, buf string, length uint32) int {
	return 0
}

func Osip_message_parse_sipfrag(sip *Osip_message_t, buf string, length uint32) int {
	return 0
}

//func Osip_message_to_str(sip *Osip_message_t, dest **) {
//
//}
