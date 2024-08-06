# go-test

**Problem statement:** Networked systems and the graph data structure are two fundamental
building blocks of our current interconnected software landscape. Using golang tools and
libraries of your choice, create an implementation of a network service that can accept a
description of an undirected non-weighted graph and return various properties about the
graph over a rest interface. The server should support more than one client connecting at
the same time.

**Command to run the applicatoin:**  go run cmd/sever/main.go

**Notes**
- The server can accept multiple requests.
  
- The application accepts Graph in two variables:
  1. nodes: This is an array of strings that denote the Network nodes
  2. edges: This is an array of strings where each element is an array of strings that denote the start and end of an edge.
 
  Sample Input:
  {"nodes":["A","B","C","D","E"],"edges":[["A","B"],["B","C"],["C","D"],["D","E"],["B","D"]]}
  
- To get the shortest path the client needs to perform GET request with the graph **id** as request parameter in the endpoint and **start** and **end** node as query parameters
  
- To delete a graph the client needs to send DELETE request with the graph **id** as request parameter
