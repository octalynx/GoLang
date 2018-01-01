type Block struct {
	Timestamp     int64
	Data          []byte
	UDID          int64
	(?)CombUMM    []
	PrevBlockHash []byte
	Hash          []byte
}





//Check sum for UDID




//Checking UDID for simultaneity




//Transaction
type Transaction struct {
	ID   []byte
	Buy  []In
	Con  []
	Sell []Out
}

//Storing transactions in chain
type Block struct {
	Timestamp     int64
	Transaction   []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

func NewBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), transaction,  prevBlockHash, []byte{}, 0}
	...
}

//Check sum for unspent
func (bc *Blockchain) FindUnspentTransaction(address string) []Transaction {
	var unspent []Transaction
	spent := make(map[string][]int)
	bci := bc.Iterator()
  
	for {
	  block := bci.Next()
  
	  for _, tx := range block.Transactions {
		txID := hex.EncodeToString(tx.ID)
  
	  Outputs:
		for out, out := range tx.Vout {
		  // Was the output spent?
		  if spent[txID] != nil {
			for _, spent := range spent[txID] {
			  if spent == out {
				continue Outputs
			  }
			}
		  }
  
		  if out.CanBeUnlockedWith(address) {
			unspent = append(unspent, *tx)
		  }
		}
  
		if tx.IsCoinbase() == false {
		  for _, in := range tx.Vin {
			if in.CanUnlockOutputWith(address) {
			  inTxID := hex.EncodeToString(in.Txid)
			  spent[inTxID] = append(spent[inTxID], in.Vout)
			}
		  }
		}
	  }
  
	  if len(block.PrevBlockHash) == 0 {
		break
	  }
	}
  
	return unspentTXs
  }




//Proof of beacon sold
type Beacon struct{
	UDID   []int64
	Major  []int64
	Minor  []int64
	ShopID []int32 //shop id, who sold beacon
	Hash   []int64 //check sum of all specs
}

func ()   //return shop id



//3Proof 

Prev. hash: 00000041662c5fc2883535dc19ba8a33ac993b535da9899e593ff98e1eda56a1
Data: Sending 5 token to 'B' for bread
UDID: B9407F30-F5F8-466E-AFF9-25556B57FE6D
CombUMM:
Hash: 00000077a856e697c69833d9effb6bdad54c730a98d674f73c0b30020cc82804

Prev. hash: 0000009ba8a33ac993b535da9841662c5fc2883598e1eda56a135dc199e593ff
Data: Receving 5 token from 'A' for bread
UDID: B9407F30-F5F8-466E-AFF9-25556B57FE6D
(?)CombUMM:  
Hash: 00000077a856e697c69833d9effb6bdad54c730a98d674f73c0b30020cc82804

Prev. hash: 0000009ba8a33ac993b535da9841662c5fc2883598e1eda56a135dc199e593ff
Data: 'A' sending 5 token ro 'B' for bread
UDID: B9407F30-F5F8-466E-AFF9-25556B57FE6D
(?)CombUMM:  
Hash: 00000077a856e697c69833d9effb6bdad54c730a98d674f73c0b30020cc82804
