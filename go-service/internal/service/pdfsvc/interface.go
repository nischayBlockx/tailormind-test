package pdfsvc

type PDFServiceInterface interface {
	GenerateStudentReport(id string) ([]byte, error)
}
