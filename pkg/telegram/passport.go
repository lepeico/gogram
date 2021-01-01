package telegram

// PassportData contains information about Telegram Passport data shared with the bot by the user.
type PassportData struct {
	// Data array with information about documents and other Telegram Passport elements
	// that was shared with the bot
	Data *[]EncryptedPassportElement `json:"data"`
	// Credentials encrypted credentials required to decrypt the data
	Credentials *EncryptedCredentials `json:"credentials"`
}

// PassportFile  represents a file uploaded to Telegram Passport.
// Currently all Telegram Passport files are in JPEG format when decrypted and don't exceed 10MB.
type PassportFile struct {
	// FileID identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`
	// FileUniqueID unique identifier for this file,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`
	// FileSize file size
	FileSize int `json:"file_size"`
	// FileDate unix time when the file was uploaded
	FileDate int `json:"file_date"`
}

// EncryptedPassportElement contains information about documents or
// other Telegram Passport elements shared with the bot by the user.
type EncryptedPassportElement struct {
	// Type element type. One of “personal_details”, “passport”, “driver_license”,
	// “identity_card”, “internal_passport”, “address”, “utility_bill”, “bank_statement”,
	// “rental_agreement”, “passport_registration”, “temporary_registration”, “phone_number”, “email”.
	Type string `json:"type"`
	// Data base64-encoded encrypted Telegram Passport element data provided by the user,
	// available for “personal_details”, “passport”, “driver_license”, “identity_card”,
	// “internal_passport” and “address” types. Can be decrypted and verified using
	// the accompanying EncryptedCredentials.
	// Optional
	Data string `json:"data"`
	// PhoneNumber user's verified phone number, available only for “phone_number” type
	// Optional
	PhoneNumber string `json:"phone_number"`
	// Email user's verified email address, available only for “email” type
	// Optional
	Email string `json:"email"`
	// Files array of encrypted files with documents provided by the user,
	// available for “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”
	// and “temporary_registration” types.
	// Files can be decrypted and verified using the accompanying EncryptedCredentials.
	// Optional
	Files *[]PassportFile `json:"files"`
	// FrontSide encrypted file with the front side of the document, provided by the user.
	// Available for “passport”, “driver_license”, “identity_card” and “internal_passport”.
	// The file can be decrypted and verified using the accompanying EncryptedCredentials.
	// Optional
	FrontSide *PassportFile `json:"front_side"`
	// ReverseSide encrypted file with the reverse side of the document, provided by the user.
	// Available for “passport”, “driver_license”, “identity_card” and “internal_passport”.
	// The file can be decrypted and verified using the accompanying EncryptedCredentials.
	// Optional
	ReverseSide *PassportFile `json:"reverse_side"`
	// Selfie encrypted file with the selfie of the user holding a document, provided by the user.
	// Available for “passport”, “driver_license”, “identity_card” and “internal_passport”.
	// The file can be decrypted and verified using the accompanying EncryptedCredentials.
	// Optional
	Selfie *PassportFile `json:"selfie"`
	// Translation array of encrypted files with translated versions of documents provided by the user.
	// Available if requested for “passport”, “driver_license”, “identity_card”, “internal_passport”,
	// “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration” and “temporary_registration” types.
	// Files can be decrypted and verified using the accompanying EncryptedCredentials.
	// Optional
	Translation *[]PassportFile `json:"translation"`
	// Hash base64-encoded element hash for using in PassportElementErrorUnspecified
	Hash string `json:"hash"`
}

// EncryptedCredentials contains data required for decrypting and authenticating EncryptedPassportElement.
// See the Telegram Passport Documentation for a complete description of the data decryption and authentication processes.
type EncryptedCredentials struct {
	// Data base64-encoded encrypted JSON-serialized data with unique user's payload,
	// data hashes and secrets required for EncryptedPassportElement decryption and authentication
	Data string `json:"data"`
	// Hash base64-encoded data hash for data authentication
	Hash string `json:"hash"`
	// Secret base64-encoded secret, encrypted with the bot's public RSA key, required for data decryption
	Secret string `json:"secret"`
}

// PassportElementErrorDataField represents an issue in one of the data fields that was provided by the user.
// The error is considered resolved when the field's value changes.
type PassportElementErrorDataField struct {
	// Source error source, must be data
	Source string `json:"source"`
	// Type the section of the user's Telegram Passport which has the error,
	// one of “personal_details”, “passport”, “driver_license”, “identity_card”, “internal_passport”, “address”
	Type string `json:"type"`
	// FieldName name of the data field which has the error
	FieldName string `json:"field_name"`
	// DataHash base64-encoded data hash
	DataHash string `json:"data_hash"`
	// Message error message
	Message string `json:"message"`
}

// PassportElementErrorFrontSide represents an issue with the front side of a document.
// The error is considered resolved when the file with the front side of the document changes.
type PassportElementErrorFrontSide struct {
	// Source error source,must be front_side
	Source string `json:"source"`
	// Type the section of the user's Telegram Passport which has the error,
	// one of “passport”, “driver_license”, “identity_card”, “internal_passport”
	Type string `json:"type"`
	// FileHash base64-encoded hash of the file with the front side of the document
	FileHash string `json:"file_hash"`
	// Message error message
	Message string `json:"message"`
}

// PassportElementErrorReverseSide represents an issue with the reverse side of a document.
// The error is considered resolved when the file with reverse side of the document changes.
type PassportElementErrorReverseSide struct {
	// Source error source,must be reverse_side
	Source string `json:"source"`
	// Type the section of the user's Telegram Passport which has the issue,
	// one of “driver_license”, “identity_card”
	Type string `json:"type"`
	// FileHash base64-encoded hash of the file with the reverse side of the document
	FileHash string `json:"file_hash"`
	// Message error message
	Message string `json:"message"`
}

// PassportElementErrorSelfie represents an issue with the selfie with a document.
// The error is considered resolved when the file with the selfie changes.
type PassportElementErrorSelfie struct {
	// Source error source,must be selfie
	Source string `json:"source"`
	// Type the section of the user's Telegram Passport which has the issue,
	// one of “passport”, “driver_license”, “identity_card”, “internal_passport”
	Type string `json:"type"`
	// FileHash base64-encoded hash of the file with the selfie
	FileHash string `json:"file_hash"`
	// Message error message
	Message string `json:"message"`
}

// PassportElementErrorFile represents an issue with a list of scans.
// The error is considered resolved when the list of files containing the scans changes.
type PassportElementErrorFile struct {
	// Source error source,must be file
	Source string `json:"source"`
	// Type the section of the user's Telegram Passport which has the issue,
	//  one of “utility_bill”, “bank_statement”, 
	// “rental_agreement”, “passport_registration”, “temporary_registration”
	Type string `json:"type"`
	// FileHash base64-encoded hash 
	FileHash string `json:"file_hash"`
	// Message error message
	Message string `json:"message"`
}

// PassportElementErrorFiles represents an issue with a list of scans.
// The error is considered resolved when the list of files containing the scans changes.
type PassportElementErrorFiles struct {
	// Source error source,must be files
	Source string `json:"source"`
	// Type the section of the user's Telegram Passport which has the issue,
	// one of “utility_bill”, “bank_statement”, 
	// “rental_agreement”, “passport_registration”, “temporary_registration”
	Type string `json:"type"`
	// FilesHash list of base64-encoded file hashes 
	FilesHash []string `json:"files_hash"`
	// Message error message
	Message string `json:"message"`
}

// PassportElementErrorTranslationFile representsan issue with one of the files that constitute the translation of a document. 
// The error is considered resolved when the file changes.
type PassportElementErrorTranslationFile struct {
	// Source error source,must be translation_file
	Source string `json:"source"`
	// Type the section of the user's Telegram Passport which has the issue,
	//  one of “passport”, “driver_license”, “identity_card”, 
	// “internal_passport”, “utility_bill”, “bank_statement”,
	//  “rental_agreement”, “passport_registration”, “temporary_registration”
	Type string `json:"type"`
	// FilesHash base64-encoded file hash 
	FileHash string `json:"files_hash"`
	// Message error message
	Message string `json:"message"`
}

// PassportElementErrorTranslationFiles represents an issue with the translated version of a document. 
// The error is considered resolved when a file with the document translation change.
type PassportElementErrorTranslationFiles struct {
	// Source error source,must be translation_files
	Source string `json:"source"`
	// Type the section of the user's Telegram Passport which has the issue,
	//  one of “passport”, “driver_license”, “identity_card”, 
	// “internal_passport”, “utility_bill”, “bank_statement”,
	//  “rental_agreement”, “passport_registration”, “temporary_registration”
	Type string `json:"type"`
	// FilesHash list of base64-encoded file hashes 
	FileHash string `json:"files_hash"`
	// Message error message
	Message string `json:"message"`
}

// PassportElementErrorUnspecified represents an issue in an unspecified place. 
// The error is considered resolved when new data is added.
type PassportElementErrorUnspecified struct {
	// Source error source,must be unspecified
	Source string `json:"source"`
	// Type the section of the user's Telegram Passport which has the issue
	Type string `json:"type"`
	// FilesHash  base64-encoded element hashes 
	FileHash string `json:"files_hash"`
	// Message error message
	Message string `json:"message"`
}