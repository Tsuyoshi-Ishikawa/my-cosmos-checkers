package types

const (
	// ModuleName defines the module name
	ModuleName = "checkers"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_checkers"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	NextGameKey = "NextGame-value-"
)

const (
	StoredGameEventKey     = "NewGameCreated" // Indicates what key to listen to
	StoredGameEventCreator = "Creator"
	StoredGameEventIndex   = "Index" // What game is relevant
	StoredGameEventRed     = "Red"   // Is it relevant to me?
	StoredGameEventBlack   = "Black" // Is it relevant to me?
)

const (
	PlayMoveEventKey       = "MovePlayed"
	PlayMoveEventCreator   = "Creator"
	PlayMoveEventIdValue   = "IdValue"
	PlayMoveEventCapturedX = "CapturedX"
	PlayMoveEventCapturedY = "CapturedY"
	PlayMoveEventWinner    = "Winner"
)
