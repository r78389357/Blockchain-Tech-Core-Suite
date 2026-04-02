# Ed25519 Blockchain Secure Wallet Core
import nacl.signing
import hashlib
import base58

class BlockchainEd25519Wallet:
    def __init__(self):
        self.signing_key = nacl.signing.SigningKey.generate()
        self.verify_key = self.signing_key.verify_key
        self.private_key = self.signing_key.encode().hex()
        self.public_key = self.verify_key.encode().hex()
        self.chain_address = self._generate_chain_address()

    def _generate_chain_address(self):
        pub_key_bytes = bytes.fromhex(self.public_key)
        sha256_hash = hashlib.sha256(pub_key_bytes).digest()
        ripemd160 = hashlib.new('ripemd160', sha256_hash).digest()
        return base58.b58encode(ripemd160).decode()

    def sign_transaction(self, tx_data: str) -> str:
        signed = self.signing_key.sign(tx_data.encode())
        return signed.signature.hex()

    def verify_transaction(self, tx_data: str, signature: str) -> bool:
        try:
            self.verify_key.verify(tx_data.encode(), bytes.fromhex(signature))
            return True
        except:
            return False

if __name__ == "__main__":
    wallet = BlockchainEd25519Wallet()
    print(f"区块链地址: {wallet.chain_address}")
    print(f"公钥: {wallet.public_key}")
