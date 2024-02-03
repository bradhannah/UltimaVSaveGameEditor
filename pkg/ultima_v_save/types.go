package ultima_v_save

type StartingMemoryAddressUb uint16
type StartingMemoryAddressU16 uint16

const (
	ubPartyMembers StartingMemoryAddressUb = 0x2B5
	ubActivePlayer                         = 0x2D5
)

const (
	u16CurrentYear StartingMemoryAddressU16 = 0x2CE
)

type CharacterClass struct {
	identifier   byte
	FriendlyName string
}

type CharacterStats struct {
	Name      string
	Class     CharacterClass
	CurrentMp byte
	MaximumMp byte

	CurrentHp byte
	MaximumHp byte

	ExperiencePoints uint16
}
