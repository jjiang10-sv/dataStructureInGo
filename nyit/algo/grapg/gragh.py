# import networkx as nx
# G = nx.Graph()
# G.add_edge("A","B", weight=4)
# G.add_edge("A","C", weight=2)
# G.add_edge("B","D", weight=12)
# #G.add_edge("A","E", weight=3)
# G.add_edge("C","E", weight=5)
# shortest_path_a_e = nx.shortest_path(G, "A", "E", weight="weight")
# print(f"shortest path from A to E is {shortest_path_a_e}")
# nx.draw(G, with_labels = True)

import networkx as nx
DG = nx.DiGraph()
DG.add_edges_from([(1,2),(2,3),(3,4),(4,5),(5,2),(4,6)])
print(f"the out edges of node 4 are {DG.out_edges(4)}")
print(f"the in degree of node 2 are {DG.in_degree(2)}")
print(f"the successor of node 4 are {list(DG.successors(4))}")
print(f"the predecessors of node 2 are {list(DG.predecessors(2))}")
nx.draw(DG, with_labels=True)

class undirectGraphList:
    def __init__(self) -> None:
        # use a dictionary to store the vertext and adjacent list of this vertex
        self.vertex_list = {}
    
    def add_vertex(self, node):
        if node in self.vertex_list:
            print("the node already exist", node)
        else:
            self.vertex_list[node] = []
    
    def add_edge(self, node1, node2):
        if node1 in self.vertex_list and node2 in self.vertex_list:
            # append edge to the list if adding a edge
            self.vertex_list[node1].append(node2)
            self.vertex_list[node2].append(node1)
        else:
            print("both the edge nodes shall be in the graph")

    def delete_vertex(self, node):
        if node not in self.vertex_list:
            print("can not delete node that is not in the graph")
        to_delete_list = self.vertex_list[node]
        del self.vertex_list[node]
        # check and loop the vertex in the neighboring nodes of the delete vertex
        for check_node in to_delete_list:
            tocheck_nodes_list = self.vertex_list[check_node]
            for i in range(0,len(tocheck_nodes_list)):
                if tocheck_nodes_list[i] == node:
                    # remove the vertex from the adjacent list
                    self.vertex_list[check_node] = tocheck_nodes_list[:i] + tocheck_nodes_list[i+1:]
                    break
    
    def delete_edge(self, node1, node2):
        if node1 in self.vertex_list and node2 in self.vertex_list:
            for node in [node1, node2]:
                # remove the node from both vertex's adjacent list
                node_list = self.vertex_list[node]
                for i in range(0, len(node_list)):
                    if node_list[i] == node:
                        self.vertex_list[node] =  node_list[:i] + node_list[i+1:]

        else:
            print("both the edge nodes shall be in the graph")

# worst case O(n*n)
class undirectGraphMatrix:
    def __init__(self, max_node_number) -> None:
        # max nodes
        self.max_node_number = max_node_number
        self.actual_nodes_number = 0
        # record the vertexes in a array
        self.vertexes = []
        # map the vertex to matrix table index
        self.vertexes_index = {}
        # init the matrix table. 0 means there is no edge between the node
        self.vertex_matrix = [ [0] * max_node_number for _ in range(max_node_number)] 
    
    def add_vertex(self, node):
        # check whether the node is already exist
        if node in self.vertexes_index:
            print("the node already exist", node)
        # check whether the length of vertexes(node list) reach the max
        elif self.actual_nodes_number == self.max_node_number:
            print("can not add for it is exceeding the max node number ", self.max_node_number)
        else:
            # no need to make change to vetexes matrix table for it is inited with the max_node_number
            self.vertexes.append(node)
            self.vertexes_index[node] = self.actual_nodes_number
            self.actual_nodes_number += 1
            
    
    def add_edge(self, node1, node2):
        if not node1 in self.vertexes_index or not node2 in self.vertexes_index:
            print("both the edge nodes shall be in the graph")
            return
        node_index1, node_index2 = self.vertexes_index[node1], self.vertexes_index[node2]
        # update teh matrix val to 1 to represent there is edge between these two nodes
        self.vertex_matrix[node_index1][node_index2] = 1
        self.vertex_matrix[node_index2][node_index1] = 1
        
            

    def delete_vertex(self, node):
        if node not in self.vertexes_index:
            print("can not delete node that is not in the graph")
        # get the index of to-delete node
        to_delete_node_index = self.vertexes_index[node]
        # update the vertexes by removing the node
        self.vertexes = self.vertexes[:to_delete_node_index] + self.vertexes[to_delete_node_index+1:]
        # remove the row of to_delete_node_index by coping over the next row till the actual_nodes_number
        for index in range(0, self.actual_nodes_number):
            # set all cols in to_delete_node_index as 0 to remove the edge
            self.vertex_matrix[index][to_delete_node_index] = 0
            # coping over the next row till the actual_nodes_number
            if index >= to_delete_node_index:
                self.vertex_matrix[index] = self.vertex_matrix[index+1]
        #set the actual_nodes_number row as all zero
        self.vertex_matrix[self.actual_nodes_number] = [0]*self.max_node_number
        # decrease the actual_nodes_number by one
        self.actual_nodes_number -=1

    def delete_edge(self, node1, node2):
        if node1 in self.vertexes_index and node2 in self.vertexes_index:
            to_delete_node_index1,to_delete_node_index2 = self.vertexes_index[node1],self.vertexes_index[node2]
            self.vertex_matrix[to_delete_node_index1][to_delete_node_index2] = 0
            self.vertex_matrix[to_delete_node_index2][to_delete_node_index1] = 0
        else:
            print("both the edge nodes shall be in the graph")