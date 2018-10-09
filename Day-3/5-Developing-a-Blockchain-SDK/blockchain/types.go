package blockchain

type Hashable interface {
	Hash() ([]byte, error)
}
