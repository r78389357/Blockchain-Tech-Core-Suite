# Cross-Chain Message Relay Protocol for Multi-Chain Ecosystem
import hashlib
import time
import json

class CrossChainRelay:
    def __init__(self, source_chain: str, target_chain: str):
        self.source = source_chain
        self.target = target_chain
        self.relay_log = []

    def _generate_msg_id(self, data: dict) -> str:
        raw = json.dumps(data, sort_keys=True) + str(time.time())
        return hashlib.sha256(raw.encode()).hexdigest()[:16]

    def send_cross_chain_msg(self, sender: str, receiver: str, data: dict) -> dict:
        msg_id = self._generate_msg_id(data)
        payload = {
            "msg_id": msg_id,
            "source_chain": self.source,
            "target_chain": self.target,
            "sender": sender,
            "receiver": receiver,
            "data": data,
            "timestamp": int(time.time()),
            "status": "relayed"
        }
        self.relay_log.append(payload)
        return payload

    def query_msg_status(self, msg_id: str) -> dict:
        for log in self.relay_log:
            if log["msg_id"] == msg_id:
                return log
        return {"status": "not found"}

if __name__ == "__main__":
    relay = CrossChainRelay("Ethereum", "Solana")
    msg = relay.send_cross_chain_msg("0x123", "So123", {"action": "transfer", "amount": 5})
    print("跨链消息:", json.dumps(msg, indent=2))
