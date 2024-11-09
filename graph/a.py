# Initialize graph with infinity representing no direct path
INF = float('inf')

# Example graph as an adjacency matrix
# Graph layout:
# A -> B = 3, A -> D = 7
# B -> A = 8, B -> C = 2
# C -> D = 1
# D -> A = 2
graph = [
    [0, 3, INF, 7],
    [8, 0, 2, INF],
    [INF, INF, 0, 1],
    [2, INF, INF, 0]
]

# Floyd-Warshall algorithm
def floyd_warshall(graph):
    # Number of vertices
    V = len(graph)
    
    # Distance matrix to store shortest paths
    dist = [row[:] for row in graph]  # Copy the initial distances from graph

    # Update distances using intermediate vertices
    for k in range(V):
        for i in range(V):
            for j in range(V):
                # Check if going through vertex k is shorter
                if dist[i][j] > dist[i][k] + dist[k][j]:
                    dist[i][j] = dist[i][k] + dist[k][j]

    # Print the result
    print("Shortest distances between every pair of vertices:")
    for i in range(V):
        for j in range(V):
            if dist[i][j] == INF:
                print("INF", end="\t")
            else:
                print(dist[i][j], end="\t")
        print()

# Run the algorithm
floyd_warshall(graph)