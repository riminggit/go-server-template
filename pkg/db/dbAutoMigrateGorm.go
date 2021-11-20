package DB

import (
	_ "github.com/go-sql-driver/mysql"
	"go-server-template/model/classify"
	"go-server-template/model/company"
	"go-server-template/model/knowledge"
	"go-server-template/model/tag"
	"go-server-template/model/topic"
	"go-server-template/model/type"
	"go-server-template/model/user"
	"gorm.io/gorm"
)

func AutoMigrateDBGorm(DB *gorm.DB) {
	DB.Table("user").AutoMigrate(&userModel.User{})
	DB.Table("user_add_topic").AutoMigrate(&userModel.UserAddTopic{})
	DB.Table("user_answer_topic_record").AutoMigrate(&userModel.UserAnswerTopicRecord{})
	DB.Table("user_collect_topic").AutoMigrate(&userModel.UserCollectTopic{})
	DB.Table("user_experience").AutoMigrate(&userModel.UserExperience{})
	DB.Table("user_feedback").AutoMigrate(&userModel.UserFeedback{})
	DB.Table("user_goals").AutoMigrate(&userModel.UserGoals{})
	DB.Table("user_interview").AutoMigrate(&userModel.UserInterview{})
	DB.Table("user_login_record").AutoMigrate(&userModel.UserLoginRecord{})
	DB.Table("user_mistake_record").AutoMigrate(&userModel.UserMistakeRecord{})
	DB.Table("user_test_record").AutoMigrate(&userModel.UserTestRecord{})
	DB.Table("user_topic_read").AutoMigrate(&userModel.UserTopicRead{})
	DB.Table("type").AutoMigrate(&typeModel.Type{})
	DB.Table("topic").AutoMigrate(&topicModel.Topic{})
	DB.Table("topic_classify").AutoMigrate(&topicModel.TopicClassify{})
	DB.Table("topic_company").AutoMigrate(&topicModel.TopicCompany{})
	DB.Table("topic_tag").AutoMigrate(&topicModel.TopicTag{})
	DB.Table("topic_type").AutoMigrate(&topicModel.TopicType{})
	DB.Table("topic_set").AutoMigrate(&topicModel.TopicSet{})
	DB.Table("tag").AutoMigrate(&tagModel.Tag{})
	DB.Table("classify").AutoMigrate(&classifyModel.Classify{})
	DB.Table("company").AutoMigrate(&companyModel.Company{})
	DB.Table("knowledge").AutoMigrate(&knowledgeModel.Knowledge{})
	DB.Table("knowledge_tag").AutoMigrate(&knowledgeModel.KnowledgeTag{})
	DB.Table("knowledge_calssify").AutoMigrate(&knowledgeModel.KnowledgeClassify{})
	DB.Table("knowledge_topic").AutoMigrate(&knowledgeModel.KnowledgeTopic{})
	DB.Table("knowledge_type").AutoMigrate(&knowledgeModel.KnowledgeType{})

	// DB.Create(&userModel.User{Email: "123456789"})
}
