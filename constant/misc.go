package constant

type ReservationCategory string

const (
	RC_ST          ReservationCategory = "ST"
	RC_SC          ReservationCategory = "SC"
	RC_GEN         ReservationCategory = "GEN"
	RC_GEN_EWS     ReservationCategory = "GEN-EWS"
	RC_GEN_PWD     ReservationCategory = "GEN-PwD"
	RC_GEN_EWS_PWD ReservationCategory = "GEN-EWS-PwD"
	RC_OBC_NCL     ReservationCategory = "OBC-NCL"
	RC_OBC_PWD     ReservationCategory = "OBC-NCL-PwD"
)
