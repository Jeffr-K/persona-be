// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ApplicationColumns holds the columns for the "Application" table.
	ApplicationColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// ApplicationTable holds the schema information for the "Application" table.
	ApplicationTable = &schema.Table{
		Name:       "Application",
		Columns:    ApplicationColumns,
		PrimaryKey: []*schema.Column{ApplicationColumns[0]},
	}
	// AssignmentColumns holds the columns for the "Assignment" table.
	AssignmentColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "type", Type: field.TypeString},
	}
	// AssignmentTable holds the schema information for the "Assignment" table.
	AssignmentTable = &schema.Table{
		Name:       "Assignment",
		Columns:    AssignmentColumns,
		PrimaryKey: []*schema.Column{AssignmentColumns[0]},
	}
	// BadgeColumns holds the columns for the "Badge" table.
	BadgeColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// BadgeTable holds the schema information for the "Badge" table.
	BadgeTable = &schema.Table{
		Name:       "Badge",
		Columns:    BadgeColumns,
		PrimaryKey: []*schema.Column{BadgeColumns[0]},
	}
	// BookmarkColumns holds the columns for the "Bookmark" table.
	BookmarkColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// BookmarkTable holds the schema information for the "Bookmark" table.
	BookmarkTable = &schema.Table{
		Name:       "Bookmark",
		Columns:    BookmarkColumns,
		PrimaryKey: []*schema.Column{BookmarkColumns[0]},
	}
	// CommentColumns holds the columns for the "Comment" table.
	CommentColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// CommentTable holds the schema information for the "Comment" table.
	CommentTable = &schema.Table{
		Name:       "Comment",
		Columns:    CommentColumns,
		PrimaryKey: []*schema.Column{CommentColumns[0]},
	}
	// CompanyColumns holds the columns for the "Company" table.
	CompanyColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// CompanyTable holds the schema information for the "Company" table.
	CompanyTable = &schema.Table{
		Name:       "Company",
		Columns:    CompanyColumns,
		PrimaryKey: []*schema.Column{CompanyColumns[0]},
	}
	// CoverLetterColumns holds the columns for the "CoverLetter" table.
	CoverLetterColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// CoverLetterTable holds the schema information for the "CoverLetter" table.
	CoverLetterTable = &schema.Table{
		Name:       "CoverLetter",
		Columns:    CoverLetterColumns,
		PrimaryKey: []*schema.Column{CoverLetterColumns[0]},
	}
	// ExperienceColumns holds the columns for the "Experience" table.
	ExperienceColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// ExperienceTable holds the schema information for the "Experience" table.
	ExperienceTable = &schema.Table{
		Name:       "Experience",
		Columns:    ExperienceColumns,
		PrimaryKey: []*schema.Column{ExperienceColumns[0]},
	}
	// FollowColumns holds the columns for the "Follow" table.
	FollowColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// FollowTable holds the schema information for the "Follow" table.
	FollowTable = &schema.Table{
		Name:       "Follow",
		Columns:    FollowColumns,
		PrimaryKey: []*schema.Column{FollowColumns[0]},
	}
	// ImageColumns holds the columns for the "Image" table.
	ImageColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// ImageTable holds the schema information for the "Image" table.
	ImageTable = &schema.Table{
		Name:       "Image",
		Columns:    ImageColumns,
		PrimaryKey: []*schema.Column{ImageColumns[0]},
	}
	// JobColumns holds the columns for the "Job" table.
	JobColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// JobTable holds the schema information for the "Job" table.
	JobTable = &schema.Table{
		Name:       "Job",
		Columns:    JobColumns,
		PrimaryKey: []*schema.Column{JobColumns[0]},
	}
	// LeetcodeColumns holds the columns for the "Leetcode" table.
	LeetcodeColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// LeetcodeTable holds the schema information for the "Leetcode" table.
	LeetcodeTable = &schema.Table{
		Name:       "Leetcode",
		Columns:    LeetcodeColumns,
		PrimaryKey: []*schema.Column{LeetcodeColumns[0]},
	}
	// LikeColumns holds the columns for the "Like" table.
	LikeColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// LikeTable holds the schema information for the "Like" table.
	LikeTable = &schema.Table{
		Name:       "Like",
		Columns:    LikeColumns,
		PrimaryKey: []*schema.Column{LikeColumns[0]},
	}
	// LocationColumns holds the columns for the "Location" table.
	LocationColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// LocationTable holds the schema information for the "Location" table.
	LocationTable = &schema.Table{
		Name:       "Location",
		Columns:    LocationColumns,
		PrimaryKey: []*schema.Column{LocationColumns[0]},
	}
	// LogoColumns holds the columns for the "Logo" table.
	LogoColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// LogoTable holds the schema information for the "Logo" table.
	LogoTable = &schema.Table{
		Name:       "Logo",
		Columns:    LogoColumns,
		PrimaryKey: []*schema.Column{LogoColumns[0]},
	}
	// NamecardColumns holds the columns for the "Namecard" table.
	NamecardColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_schema_namecard", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// NamecardTable holds the schema information for the "Namecard" table.
	NamecardTable = &schema.Table{
		Name:       "Namecard",
		Columns:    NamecardColumns,
		PrimaryKey: []*schema.Column{NamecardColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "Namecard_User_namecard",
				Columns:    []*schema.Column{NamecardColumns[1]},
				RefColumns: []*schema.Column{UserColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// NotificationColumns holds the columns for the "Notification" table.
	NotificationColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// NotificationTable holds the schema information for the "Notification" table.
	NotificationTable = &schema.Table{
		Name:       "Notification",
		Columns:    NotificationColumns,
		PrimaryKey: []*schema.Column{NotificationColumns[0]},
	}
	// PaymentColumns holds the columns for the "Payment" table.
	PaymentColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// PaymentTable holds the schema information for the "Payment" table.
	PaymentTable = &schema.Table{
		Name:       "Payment",
		Columns:    PaymentColumns,
		PrimaryKey: []*schema.Column{PaymentColumns[0]},
	}
	// PersonalizationColumns holds the columns for the "Personalization" table.
	PersonalizationColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_schema_personalization", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// PersonalizationTable holds the schema information for the "Personalization" table.
	PersonalizationTable = &schema.Table{
		Name:       "Personalization",
		Columns:    PersonalizationColumns,
		PrimaryKey: []*schema.Column{PersonalizationColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "Personalization_User_personalization",
				Columns:    []*schema.Column{PersonalizationColumns[1]},
				RefColumns: []*schema.Column{UserColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PhoneColumns holds the columns for the "Phone" table.
	PhoneColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "profile_schema_phone", Type: field.TypeInt, Unique: true},
	}
	// PhoneTable holds the schema information for the "Phone" table.
	PhoneTable = &schema.Table{
		Name:       "Phone",
		Columns:    PhoneColumns,
		PrimaryKey: []*schema.Column{PhoneColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "Phone_Profile_phone",
				Columns:    []*schema.Column{PhoneColumns[1]},
				RefColumns: []*schema.Column{ProfileColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PhotoColumns holds the columns for the "photo" table.
	PhotoColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "profile_schema_photo", Type: field.TypeInt, Unique: true},
	}
	// PhotoTable holds the schema information for the "photo" table.
	PhotoTable = &schema.Table{
		Name:       "photo",
		Columns:    PhotoColumns,
		PrimaryKey: []*schema.Column{PhotoColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "photo_Profile_photo",
				Columns:    []*schema.Column{PhotoColumns[1]},
				RefColumns: []*schema.Column{ProfileColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PoolColumns holds the columns for the "Pool" table.
	PoolColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// PoolTable holds the schema information for the "Pool" table.
	PoolTable = &schema.Table{
		Name:       "Pool",
		Columns:    PoolColumns,
		PrimaryKey: []*schema.Column{PoolColumns[0]},
	}
	// PortfolioColumns holds the columns for the "Portfolio" table.
	PortfolioColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// PortfolioTable holds the schema information for the "Portfolio" table.
	PortfolioTable = &schema.Table{
		Name:       "Portfolio",
		Columns:    PortfolioColumns,
		PrimaryKey: []*schema.Column{PortfolioColumns[0]},
	}
	// PositionColumns holds the columns for the "Position" table.
	PositionColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// PositionTable holds the schema information for the "Position" table.
	PositionTable = &schema.Table{
		Name:       "Position",
		Columns:    PositionColumns,
		PrimaryKey: []*schema.Column{PositionColumns[0]},
	}
	// ProfileColumns holds the columns for the "Profile" table.
	ProfileColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_schema_profile", Type: field.TypeInt, Unique: true},
	}
	// ProfileTable holds the schema information for the "Profile" table.
	ProfileTable = &schema.Table{
		Name:       "Profile",
		Columns:    ProfileColumns,
		PrimaryKey: []*schema.Column{ProfileColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "Profile_User_profile",
				Columns:    []*schema.Column{ProfileColumns[1]},
				RefColumns: []*schema.Column{UserColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ReferrerColumns holds the columns for the "Referrer" table.
	ReferrerColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_schema_referrer", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// ReferrerTable holds the schema information for the "Referrer" table.
	ReferrerTable = &schema.Table{
		Name:       "Referrer",
		Columns:    ReferrerColumns,
		PrimaryKey: []*schema.Column{ReferrerColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "Referrer_User_referrer",
				Columns:    []*schema.Column{ReferrerColumns[1]},
				RefColumns: []*schema.Column{UserColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ResumeColumns holds the columns for the "Resume" table.
	ResumeColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// ResumeTable holds the schema information for the "Resume" table.
	ResumeTable = &schema.Table{
		Name:       "Resume",
		Columns:    ResumeColumns,
		PrimaryKey: []*schema.Column{ResumeColumns[0]},
	}
	// RoleColumns holds the columns for the "Role" table.
	RoleColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_schema_roles", Type: field.TypeInt, Nullable: true},
	}
	// RoleTable holds the schema information for the "Role" table.
	RoleTable = &schema.Table{
		Name:       "Role",
		Columns:    RoleColumns,
		PrimaryKey: []*schema.Column{RoleColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "Role_User_roles",
				Columns:    []*schema.Column{RoleColumns[1]},
				RefColumns: []*schema.Column{UserColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SkillColumns holds the columns for the "Skill" table.
	SkillColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// SkillTable holds the schema information for the "Skill" table.
	SkillTable = &schema.Table{
		Name:       "Skill",
		Columns:    SkillColumns,
		PrimaryKey: []*schema.Column{SkillColumns[0]},
	}
	// SubscribeColumns holds the columns for the "Subscribe" table.
	SubscribeColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// SubscribeTable holds the schema information for the "Subscribe" table.
	SubscribeTable = &schema.Table{
		Name:       "Subscribe",
		Columns:    SubscribeColumns,
		PrimaryKey: []*schema.Column{SubscribeColumns[0]},
	}
	// ThreadColumns holds the columns for the "Thread" table.
	ThreadColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// ThreadTable holds the schema information for the "Thread" table.
	ThreadTable = &schema.Table{
		Name:       "Thread",
		Columns:    ThreadColumns,
		PrimaryKey: []*schema.Column{ThreadColumns[0]},
	}
	// UserColumns holds the columns for the "User" table.
	UserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uuid", Type: field.TypeUUID},
		{Name: "username", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_schema_follow", Type: field.TypeInt, Nullable: true},
	}
	// UserTable holds the schema information for the "User" table.
	UserTable = &schema.Table{
		Name:       "User",
		Columns:    UserColumns,
		PrimaryKey: []*schema.Column{UserColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "User_Follow_follow",
				Columns:    []*schema.Column{UserColumns[7]},
				RefColumns: []*schema.Column{FollowColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ApplicationTable,
		AssignmentTable,
		BadgeTable,
		BookmarkTable,
		CommentTable,
		CompanyTable,
		CoverLetterTable,
		ExperienceTable,
		FollowTable,
		ImageTable,
		JobTable,
		LeetcodeTable,
		LikeTable,
		LocationTable,
		LogoTable,
		NamecardTable,
		NotificationTable,
		PaymentTable,
		PersonalizationTable,
		PhoneTable,
		PhotoTable,
		PoolTable,
		PortfolioTable,
		PositionTable,
		ProfileTable,
		ReferrerTable,
		ResumeTable,
		RoleTable,
		SkillTable,
		SubscribeTable,
		ThreadTable,
		UserTable,
	}
)

func init() {
	ApplicationTable.Annotation = &entsql.Annotation{
		Table: "Application",
	}
	AssignmentTable.Annotation = &entsql.Annotation{
		Table: "Assignment",
	}
	BadgeTable.Annotation = &entsql.Annotation{
		Table: "Badge",
	}
	BookmarkTable.Annotation = &entsql.Annotation{
		Table: "Bookmark",
	}
	CommentTable.Annotation = &entsql.Annotation{
		Table: "Comment",
	}
	CompanyTable.Annotation = &entsql.Annotation{
		Table: "Company",
	}
	CoverLetterTable.Annotation = &entsql.Annotation{
		Table: "CoverLetter",
	}
	ExperienceTable.Annotation = &entsql.Annotation{
		Table: "Experience",
	}
	FollowTable.Annotation = &entsql.Annotation{
		Table: "Follow",
	}
	ImageTable.Annotation = &entsql.Annotation{
		Table: "Image",
	}
	JobTable.Annotation = &entsql.Annotation{
		Table: "Job",
	}
	LeetcodeTable.Annotation = &entsql.Annotation{
		Table: "Leetcode",
	}
	LikeTable.Annotation = &entsql.Annotation{
		Table: "Like",
	}
	LocationTable.Annotation = &entsql.Annotation{
		Table: "Location",
	}
	LogoTable.Annotation = &entsql.Annotation{
		Table: "Logo",
	}
	NamecardTable.ForeignKeys[0].RefTable = UserTable
	NamecardTable.Annotation = &entsql.Annotation{
		Table: "Namecard",
	}
	NotificationTable.Annotation = &entsql.Annotation{
		Table: "Notification",
	}
	PaymentTable.Annotation = &entsql.Annotation{
		Table: "Payment",
	}
	PersonalizationTable.ForeignKeys[0].RefTable = UserTable
	PersonalizationTable.Annotation = &entsql.Annotation{
		Table: "Personalization",
	}
	PhoneTable.ForeignKeys[0].RefTable = ProfileTable
	PhoneTable.Annotation = &entsql.Annotation{
		Table: "Phone",
	}
	PhotoTable.ForeignKeys[0].RefTable = ProfileTable
	PhotoTable.Annotation = &entsql.Annotation{
		Table: "photo",
	}
	PoolTable.Annotation = &entsql.Annotation{
		Table: "Pool",
	}
	PortfolioTable.Annotation = &entsql.Annotation{
		Table: "Portfolio",
	}
	PositionTable.Annotation = &entsql.Annotation{
		Table: "Position",
	}
	ProfileTable.ForeignKeys[0].RefTable = UserTable
	ProfileTable.Annotation = &entsql.Annotation{
		Table: "Profile",
	}
	ReferrerTable.ForeignKeys[0].RefTable = UserTable
	ReferrerTable.Annotation = &entsql.Annotation{
		Table: "Referrer",
	}
	ResumeTable.Annotation = &entsql.Annotation{
		Table: "Resume",
	}
	RoleTable.ForeignKeys[0].RefTable = UserTable
	RoleTable.Annotation = &entsql.Annotation{
		Table: "Role",
	}
	SkillTable.Annotation = &entsql.Annotation{
		Table: "Skill",
	}
	SubscribeTable.Annotation = &entsql.Annotation{
		Table: "Subscribe",
	}
	ThreadTable.Annotation = &entsql.Annotation{
		Table: "Thread",
	}
	UserTable.ForeignKeys[0].RefTable = FollowTable
	UserTable.Annotation = &entsql.Annotation{
		Table: "User",
	}
}
