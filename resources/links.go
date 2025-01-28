package resources

const (
	LinkUserStorageBase      = "http://localhost:8002/users-storage/v1"
	LinkOrganizationBase     = "http://localhost:8003/org-storage/v1"
	LinkInitiativeBase       = "http://localhost:8004/initiative-storage/v1"
	LinkChatStorageBase      = "http://localhost:8008/chat-storage/v1"
	LinkReactionsStorageBase = "http://localhost:8009/reactions-storage/v1"

	LinkGetUser       = "/public/users/"
	LinkGetOrg        = "/public/organization/"
	LinkGetChat       = "/public/chat/"
	LinkGetReactions  = "/public/reactions/"
	LinkGetInitiative = "/public/initiatives/"
	LinkPoints        = "/points"

	LinkPrivateUser       = "/private/users/"
	LinkPrivatePlan       = "/private/points/"
	LinkPrivateOrg        = "/private/organization/"
	LinkPrivateChat       = "/private/chat/"
	LinkPrivateInitiative = "/private/initiatives/"
	LinkPrivateReactions  = "/private/reactions/"

	LinkParticipants = "/participants"

	Likes   = "likes/"
	Reposts = "reposts/"
	Reports = "reports/"

	Pagination10EndLink = "?page[size]=10&page[number]=10"
)
