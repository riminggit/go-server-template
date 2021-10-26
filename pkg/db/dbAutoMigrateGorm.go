package DB

import (
	_ "github.com/go-sql-driver/mysql"
	"go-server-template/model/classify"
	"go-server-template/model/company"
	"go-server-template/model/tag"
	"go-server-template/model/topic"
	"go-server-template/model/type"
	"go-server-template/model/user"
	"gorm.io/gorm"
)

func AutoMigrateDBGorm(DB *gorm.DB) {

	DB.AutoMigrate(
		&userModel.User{},
		&userModel.UserAddTopic{},
		&userModel.UserAnswerTopicRecord{},
		&userModel.UserCollectTopic{},
		&userModel.UserExperience{},
		&userModel.UserFeedback{},
		&userModel.UserGoals{},
		&userModel.UserInterview{},
		&userModel.UserLoginRecord{},
		&userModel.UserMistakeRecord{},
		&userModel.UserTestRecord{},
		&userModel.UserTopicRead{},
		&typeModel.Type{},
		&topicModel.TopicType{},
		&topicModel.Topic{},
		&topicModel.TopicTag{},
		&topicModel.TopicCompany{},
		&topicModel.TopicClassify{},
		&tagModel.Tag{},
		&companyModel.Company{},
		&classifyModel.Classify{},
	)

	// DB.Create(&userModel.User{Account: "123456789"})
}
