package errors

type ErrorMessage struct {
	EN string
	ID string
}

type ErrorMessages map[string]ErrorMessage

var (
	EM = ErrorMessages{
		"internal": ErrorMessage{
			EN: `Internal Server Error. Please Call Administrator.`,
			ID: `Terjadi Kendala Pada Server. Mohon Hubungi Administrator.`,
		},
		"notfound": ErrorMessage{
			EN: `Record Does Not Exist. Please Validate Your Input Or Contact Administrator.`,
			ID: `Data Tidak Diketemukan. Mohon Cek Kembali Masukkan Anda Atau Hubungi Administrator.`,
		},
		"badrequest": ErrorMessage{
			EN: `Invalid Input. Please Validate Your Input.`,
			ID: `Kesalahan Input. Mohon Cek Kembali Masukkan Anda.`,
		},
		"unauthorized": ErrorMessage{
			EN: `Unauthorized Access. You are not authorized to access this resource.`,
			ID: `Akses Ditolak. Anda Belum Diijinkan Untuk Mengakses Aplikasi.`,
		},
		"uniqueconst": ErrorMessage{
			EN: `Record Has Existed and Must Be Unique. Please Validate Your Input Or Contact Administrator.`,
			ID: `Data sudah ada. Mohon Cek Kembali Masukkan Anda Atau Hubungi Administrator.`,
		},
		"refreshtokenexpired": ErrorMessage{
			EN: `Session Refresh Token Had Been Expired. Please Renew Your Session.`,
			ID: `Kunci Pembaharuan Sesi Sudah Berakhir. Mohon Perbaharui Sesi Anda.`,
		},
		"accesstokenexpired": ErrorMessage{
			EN: `Session Access Token Had Been Expired. Please Renew Your Session.`,
			ID: `Kunci Akses Sesi Sudah Berakhir. Mohon Perbaharui Sesi Anda.`,
		},
		"alreadyexist": ErrorMessage{
			EN: `Record already exist. Please Validate Your Input Or Contact Administrator.`,
			ID: `Data sudah ada. Mohon Cek Kembali Masukkan Anda Atau Hubungi Administrator.`,
		},
		"doesnotmatch": ErrorMessage{
			EN: `Some data sent do not match comparing to database application. Please re-check your data`,
			ID: `Data yang dikirimkan tidak cocok dengan yang ada di sistem basisdata. Tolong periksa kembali data Anda`,
		},
		"serviceunavailable": ErrorMessage{
			EN: "Service Unavailable. Please try again later.",
			ID: "Layanan ini tidak tersedia. Silakan dicoba kembali nanti.",
		},
	}
)

func (em ErrorMessages) Message(lang string, i string) string {
	if lang == "ID" {
		return em[i].ID
	}
	return em[i].EN
}
