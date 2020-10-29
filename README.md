## Certificate/Key Files

- `cert.pem` and `key.pem`: Simple RSA certificate and key pair. The key is not password protected.
- `cert_pass.pem` and `key_pass.pem`: RSA certificate and key pair where the key is password protected. The password is `passphrase`.
- `combined_cert.pem`: Contains `cert.pem` and `cert_pass.pem` concatenated, in that order.