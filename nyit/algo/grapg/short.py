# Initialize graph with infinity representing no direct path
INF = float('inf')

# Example graph as an adjacency matrix for 7 nodes (A to G)
# Node layout:
# A -> B = 3, A -> D = 7
# B -> C = 2
# C -> G = 5, C -> D = 1
# D -> A = 2, D -> F = 4
# E -> F = 6
# F -> G = 3
graph = [
    [0, 3, INF, 7, INF, INF, INF],   # A
    [INF, 0, 2, INF, INF, INF, INF], # B
    [INF, INF, 0, 1, INF, INF, 5],   # C
    [2, INF, INF, 0, INF, 4, INF],   # D
    [INF, INF, INF, INF, 0, 6, INF], # E
    [INF, INF, INF, INF, INF, 0, 3], # F
    [INF, INF, INF, INF, INF, INF, 0]# G
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

class UnionFind:
    def __init__(self, n):
        self.parent = list(range(n))
        self.rank = [0] * n

    def find(self, u):
        if self.parent[u] != u:
            self.parent[u] = self.find(self.parent[u])  # Path compression
        return self.parent[u]

    def union(self, u, v):
        root_u = self.find(u)
        root_v = self.find(v)
        if root_u != root_v:
            # Union by rank
            if self.rank[root_u] > self.rank[root_v]:
                self.parent[root_v] = root_u
            elif self.rank[root_u] < self.rank[root_v]:
                self.parent[root_u] = root_v
            else:
                self.parent[root_v] = root_u
                self.rank[root_u] += 1

def kruskal_mst(edges, n):
    uf = UnionFind(n)
    mst = []
    edges.sort(key=lambda x: x[2])  # Sort edges by weight

    for u, v, weight in edges:
        if uf.find(u) != uf.find(v):
            uf.union(u, v)
            mst.append((u, v, weight))

    return mst

edges = [
    (0, 1, 10),
    (0, 2, 6),
    (0, 3, 5),
    (1, 3, 15),
    (2, 3, 4)
]
n = 4  # Number of vertices

mst = kruskal_mst(edges, n)
print("Edges in Kruskal's MST:", mst)

import heapq

def prim_mst(graph, start=0):
    n = len(graph)
    mst = []
    visited = [False] * n
    min_heap = [(0, start, -1)]  # (weight, node, parent)

    while min_heap and len(mst) < n - 1:
        weight, u, parent = heapq.heappop(min_heap)
        if visited[u]:
            continue

        visited[u] = True
        if parent != -1:
            mst.append((parent, u, weight))

        for v, w in graph[u]:
            if not visited[v]:
                heapq.heappush(min_heap, (w, v, u))

    return mst

graph = {
    0: [(1, 10), (2, 6), (3, 5)],
    1: [(0, 10), (3, 15)],
    2: [(0, 6), (3, 4)],
    3: [(0, 5), (1, 15), (2, 4)]
}

mst = prim_mst(graph)
print("Edges in Prim's MST:", mst)