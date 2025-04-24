# Crypto Tools

Contains various commands for exploring various cryptographic primitives.
It Can be used as a library or as a CLI tool.

The Purpose is to learn more about certain cryptographic primitives, most notably Elliptic Curve Cryptography.

Currently, supports some operations on ECDSA. 

## ECDSA key recovery from nonce

```shell
ct ecdsa crack nonce reuse -c p256 -m "Hello" -n 93356471150125927255015609149458882240 -s 3046022100aa5367e2f9d615295fc5395e118640d428e2941db6f98d3e9ac6407dc27a332a022100c783f3bea45eb6cd19f647354fea2813d1062065451b8cebee2ff2766b16eeaa
```

Further examples are in the unit tests.
