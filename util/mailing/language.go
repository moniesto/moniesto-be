package mailing

import (
	"fmt"

	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/util/system"
)

type TemplateType struct {
	Path    string
	Subject string
}

type EmailTemplateTypes struct {
	en TemplateType
	tr TemplateType
}

var EmailTemplates map[string]EmailTemplateTypes = map[string]EmailTemplateTypes{
	"email_verification": {
		en: TemplateType{Path: "util/mailing/templates/email_verification/email_verification_en.html", Subject: "Moniesto: Email Verification"},
		tr: TemplateType{Path: "util/mailing/templates/email_verification/email_verification_tr.html", Subject: "Moniesto: E-posta Doğrulama"},
	},
	"welcoming": {
		en: TemplateType{Path: "util/mailing/templates/welcoming/welcoming_en.html", Subject: "Moniesto: Thank You for Joining Moniesto!"},
		tr: TemplateType{Path: "util/mailing/templates/welcoming/welcoming_tr.html", Subject: "Moniesto: Moniesto'ya Katıldığınız İçin Teşekkürler!"},
	},
	"password-reset": {
		en: TemplateType{Path: "util/mailing/templates/password_reset/password_reset_en.html", Subject: "Moniesto: Reset password!"},
		tr: TemplateType{Path: "util/mailing/templates/password_reset/password_reset_tr.html", Subject: "Moniesto: Şifreni sıfırla!"},
	},
	"new_post": {
		en: TemplateType{Path: "util/mailing/templates/new_post/new_post_en.html", Subject: "Moniesto: New Post!"},
		tr: TemplateType{Path: "util/mailing/templates/new_post/new_post_tr.html", Subject: "Moniesto: Yeni Gönderi!"},
	},
	"payout": {
		en: TemplateType{Path: "util/mailing/templates/payout/payout_en.html", Subject: "Moniesto: Monthly Payout"},
		tr: TemplateType{Path: "util/mailing/templates/payout/payout_tr.html", Subject: "Moniesto: Aylık Ödeme"},
	},
	"subscribe_user": {
		en: TemplateType{Path: "util/mailing/templates/subscribe/subscribe_user_en.html", Subject: "Moniesto: Subscription Started"},
		tr: TemplateType{Path: "util/mailing/templates/subscribe/subscribe_user_tr.html", Subject: "Moniesto: Aboneliğiniz Başladı"},
	},
	"subscribe_moniest": {
		en: TemplateType{Path: "util/mailing/templates/subscribe/subscribe_moniest_en.html", Subject: "Moniesto: New Subscriber!"},
		tr: TemplateType{Path: "util/mailing/templates/subscribe/subscribe_moniest_tr.html", Subject: "Moniesto: Yeni Abone!"},
	},
}

func GetTemplate(templateName string, language db.UserLanguage) (TemplateType, error) {
	EmailTemplateTypes, ok := EmailTemplates[templateName]
	if !ok {
		return TemplateType{}, fmt.Errorf("template not found: %s", templateName)
	}

	switch language {
	case db.UserLanguageEn:
		return EmailTemplateTypes.en, nil
	case db.UserLanguageTr:
		return EmailTemplateTypes.tr, nil
	case "": // empty language case
		system.LogError("language is empty, using -en- as default")
		return EmailTemplateTypes.en, nil
	default:
		return TemplateType{}, fmt.Errorf("template not found: %s with language %s", templateName, language)
	}
}
