set -x
sed -i "s/$MTT_DYNAMIC_PASSWORD/MTT_DYNAMIC_PASSWORD/g" serverGOLANG.go
sed -i "s/$MTT_DYNAMIC_MAIL_PASSWORD/MTT_DYNAMIC_MAIL_PASSWORD/g" sendmailGOLANG.go
set +x
