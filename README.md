## POW algorithm
POW algorithm that was chosen is a simple cryptographic puzzle. Client has to find a prefix that when hashed together with nonce gives a hex with `n` leading zeros.
This is a simple enough assignment that uses only CPU and can be easily tweaked to either stall the clients more or less (by tweaking difficulty). The choice
for this algorithm is its simplicity of implementation and flexibility with the time it will take clients to solve it. We don't want them to take to long because then
we completely lose the point of real time communication, but we want to stall them enough to lessen the effects of a DDOS.

## Protocol
Protocol consists of several phases:
1. Client sends server a RequestChallenge packet
2. Server responds with a nonce and a SendChallenge packet
3. Client then calculates the prefix and sends a SendSolution packet
4. Server validates the solution and sends one of the following:
    1. SolutionIncorrect packet that contains the nonce and difficulty so that client start again easily
    2. SendQuote packet that contains the words of wisdom quote

There is also an Error package type that is used for crude error communication to the client.

All messages are JSON encoded as that is the most user friendly, easy for implementation,
reusable in other languages and we can also validate it with json schema. Protobuf would be 
my encoding of choice but it's a bit harder to setup and since it's not plaintext it's a bit
less user friendly and easy to debug, you need additional tooling.
