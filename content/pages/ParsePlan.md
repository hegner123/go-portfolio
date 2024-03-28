# Parse Plan

1. open input file

2. scan input file byte by byte
 - If byte is a valid unicode character, add it to the current token.
 - If byte is a whitespace, add the current token to the token list and start a new token.
 - If byte is verticle whitespace (newline), and is proceeded by a verticle whitespace discard it.
 - If byte is a verticle whitespace, add the current token to the token list and start a new token.

3. close input file

4. return token list




