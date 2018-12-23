package main

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

// type Blockchain struct {
// 	tip []byte
// 	db  *bolt.DB
// }

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// func (bc *Blockchain) AddBlock(data string) {
// 	var lastHash []byte

// 	err := bc.db.View(func(tx *bolt.Tx) error {
// 		b := tx.Bucket([]byte(blocksBucket))
// 		lastHash = b.Get([]byte("l"))

// 		return nil
// 	})

// 	// mining newblock
// 	newBlock := NewBlock(data, lastHash)

// 	err = bc.db.Update(func(tx *bolt.Tx) error {
// 		b := tx.Bucket([]byte(blocksBucket))
// 	})
// }

func NewBlockchain() *Blockchain {
	newGenesisBlock := NewGenesisBlock()
	return &Blockchain{[]*Block{newGenesisBlock}}
}

// func NewBlockchain() *Blockchain {
// 	var tip []byte
// 	db, err := bolt.Open(dbFile, 0600, nil)

// 	err = db.Update(func(tx *bolt.Tx) error {
// 		b := tx.Bucket([]byte(blocksBucket))

// 		// cannot find any db, just create new one.
// 		if b == nil {
// 			genesis := NewGenesisBlock()
// 			b, err := tx.CreateBucket([]byte(blocksBucket))
// 			if err != nil {
// 				log.Panic(err)
// 			}

// 			// key - value database
// 			// 32-byte block-hash -> Block structure (serialized)
// 			err = b.Put(genesis.Hash, genesis.Serialize())
// 			if err != nil {
// 				log.Panic(err)
// 			}

// 			// 'l' -> the hash of the last block in a chain
// 			err = b.Put([]byte("l"), genesis.Hash)
// 			if err != nil {
// 				log.Panic(err)
// 			}
// 			tip = genesis.Hash
// 		} else {
// 			tip = b.Get([]byte("l"))
// 		}

// 		return nil
// 	})

// 	return &Blockchain{tip, db}
// }
