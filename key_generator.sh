#!/bin/bash

openssl genrsa -out private.key 4096
openssl rsa -in private.key -out public.key -pubout