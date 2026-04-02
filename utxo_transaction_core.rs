// UTXO Transaction Core for Bitcoin-like Blockchains
use sha2::{Sha256, Digest};
use hex::encode;

#[derive(Debug, Clone)]
pub struct UTXOInput {
    pub txid: String,
    pub vout: u32,
    pub signature: String,
}

#[derive(Debug, Clone)]
pub struct UTXOOutput {
    pub address: String,
    pub amount: u64,
}

#[derive(Debug)]
pub struct UTXOTransaction {
    pub inputs: Vec<UTXOInput>,
    pub outputs: Vec<UTXOOutput>,
}

impl UTXOTransaction {
    pub fn new() -> Self {
        UTXOTransaction {
            inputs: Vec::new(),
            outputs: Vec::new(),
        }
    }

    pub fn add_input(&mut self, txid: String, vout: u32, sig: String) {
        self.inputs.push(UTXOInput { txid, vout, signature: sig });
    }

    pub fn add_output(&mut self, addr: String, amt: u64) {
        self.outputs.push(UTXOOutput { address: addr, amount: amt });
    }

    pub fn get_tx_hash(&self) -> String {
        let data = format!("{:?}{:?}", self.inputs, self.outputs);
        let hash = Sha256::digest(data.as_bytes());
        encode(hash)
    }
}

fn main() {
    let mut tx = UTXOTransaction::new();
    tx.add_input("abcd1234".to_string(), 0, "sig_hex".to_string());
    tx.add_output("bc1qxyz".to_string(), 50000);
    println!("TX Hash: {}", tx.get_tx_hash());
}
