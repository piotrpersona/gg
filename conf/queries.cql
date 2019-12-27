CREATE (nAlice:User {Name: 'Alice', seed: 42})CREATE (nBridget:User {Name: 'Bridget', seed: 42})CREATE (nCharles:User {Name: 'Charles', seed: 42})CREATE (nDoug:User {Name: 'Doug'})CREATE (nMark:User {Name: 'Mark'})CREATE (nMichael:User {Name: 'Michael'})CREATE (nAlice)-[:LINK {weight: 1}]->(nBridget)CREATE (nAlice)-[:LINK {weight: 1}]->(nCharles)CREATE (nCharles)-[:LINK {weight: 1}]->(nBridget)CREATE (nAlice)-[:LOVES {weight: 5}]->(nDoug)CREATE (nMark)-[:LINK {weight: 1}]->(nDoug)CREATE (nMark)-[:LINK {weight: 1}]->(nMichael)CREATE (nMichael)-[:LINK {weight: 1}]->(nMark)

CALL algo.beta.louvain.stream('User', 'LINK', {graph: 'huge', direction: 'BOTH'}) YIELD nodeId, community, communities RETURN algo.asNode(nodeId).Name as name, community, communities ORDER BY name ASC


CALL algo.beta.louvain.stream('User', 'LINK', {graph: 'huge', direction: 'BOTH', weightProperty: 'weight'}) YIELD nodeId, community, communities RETURN algo.asNode(nodeId).Name as name, community, communities ORDER BY name ASC
