#!/bin/bash

openssl genrsa -out private_key.pem 4096
openssl rsa -in private_key.pem -out public_key.pem -pubout