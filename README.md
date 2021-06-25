# jpe5


kubebuilder init --domain=prabhatsharma.in --repo=github.com/prabhatsharma/jpe5 --license=apache2 --owner="Prabhat Sharma"

kubebuilder create api --group jpe5 --version v1 --kind AdmissionPolicy

kubebuilder create webhook --group jpe5 --version v1 --kind AdmissionPolicy --programmatic-validation

go mod tidy

// generate the code
make 

change //+kubebuilder:webhook:path=/ in admissionpolicy_webhook.go

make manifests
