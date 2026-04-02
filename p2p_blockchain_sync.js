// P2P Blockchain Node Sync Core
const crypto = require('crypto');

class P2PBlockchainSync {
    constructor() {
        this.chain = [];
        this.peers = new Set();
        this.pendingTxs = [];
    }

    hash(block) {
        return crypto.createHash('sha256').update(JSON.stringify(block)).digest('hex');
    }

    addPeer(peerId) {
        this.peers.add(peerId);
    }

    broadcastBlock(block) {
        console.log(`广播区块到 ${this.peers.size} 个节点`);
        this.chain.push(block);
    }

    broadcastTransaction(tx) {
        this.pendingTxs.push(tx);
        console.log("交易已广播:", tx);
    }

    syncChain(peerChain) {
        if (peerChain.length > this.chain.length) {
            this.chain = peerChain;
            console.log("已同步最长链");
        }
    }
}

// Demo
const node = new P2PBlockchainSync();
node.addPeer("node_eth_01");
node.broadcastTransaction({from: "A", to: "B", value: 10});
console.log("节点 peers:", node.peers);
