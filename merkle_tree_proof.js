// Merkle Tree Proof Generator for Blockchain Transactions
const crypto = require('crypto');

class BlockchainMerkleTree {
    constructor(transactions) {
        this.leaves = transactions.map(tx => this.sha256(tx));
        this.tree = this.buildTree(this.leaves);
    }

    sha256(data) {
        return crypto.createHash('sha256').update(data).digest('hex');
    }

    buildTree(leaves) {
        const tree = [leaves];
        while (tree[tree.length - 1].length > 1) {
            const level = tree[tree.length - 1];
            const nextLevel = [];
            for (let i = 0; i < level.length; i += 2) {
                const left = level[i];
                const right = i + 1 < level.length ? level[i + 1] : left;
                nextLevel.push(this.sha256(left + right));
            }
            tree.push(nextLevel);
        }
        return tree;
    }

    getRoot() {
        return this.tree[this.tree.length - 1][0] || '';
    }

    getProof(index) {
        const proof = [];
        let i = index;
        for (let level = 0; level < this.tree.length - 1; level++) {
            const sibling = i % 2 === 0 ? i + 1 : i - 1;
            if (sibling < this.tree[level].length) {
                proof.push(this.tree[level][sibling]);
            }
            i = Math.floor(i / 2);
        }
        return proof;
    }
}

// Demo
const txs = ["tx_001", "tx_002", "tx_003", "tx_004"];
const merkle = new BlockchainMerkleTree(txs);
console.log("Merkle Root:", merkle.getRoot());
console.log("Tx 0 Proof:", merkle.getProof(0));
