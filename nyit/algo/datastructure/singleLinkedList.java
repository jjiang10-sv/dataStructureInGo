class Node {
    int data;
    Node next;

    public Node (int data) {
        this.data = data;
        this.next = null;
    }
}

class SingleLinkedList {
    private Node head;
    private Node end;

    public insertAtHead(int data) {
        Node newNode = new Node(data);
        newNode.next = this.head;
        this.head = newNode;
        if this.head.next == null {
            this.end = newNode;
        }
    }
    public insertAtEnd(int data) {
        Node newNode = new Node(data);
        if this.end != null {
            this.end.next = newNode;
        } else {
            // the case of head not null but end as null not exist; so only the case of head and end as null
            this.head = newNode;
            this.end = newNode;
        }
    }
    public removeAtHead() {
        if this.head != null {
            this.head = this.head.next;
        }else {
            system.out.println("the linked list has no items to remove")
        }
    }

    public remoteAtEnd() {
        if this.end != null {
            // this.head must exist
            this.end = null;
            Node temp = this.head
            while (temp.next != null) {
                temp = temp.next
            }
            this.end == temp;
        }
    }

    public removeData(int data) {
        if this.head == null {
            return;
        }
        if this.head.data == data {
            this.head = this.head.next;
            return;
        }
        Node temp = this.head
        while (temp.next != null) and temp.next.data != data  {
            temp = temp.next
            
        }
        if temp.next.next != null {
            temp.next = temp.next.next
        }
        return;
    }
}
class DoubleListedListNode {
    int data;
    DoubleLinkedList left;
    DoubleLinkedList right;

    public DoubleListedListNode (int data){
        this.data = data;
        this.left = null;
        this.right = null;
    }
}
class DoubleLinkedList {
    private DoubleListedListNode head;
    private DoubleListedListNode end;

    public insertAtHead(int data) {
        Node newNode = new Node(data);
        if this.head == null {
            this.head = newNode;
            this.end = null;
        }else {
            newNode.left = this.head;
            if this.end == null {
                this.end = this.head
                this.head = newNode
            }else{
                this.head = newNode
            }
        }
    }
    public insertAtEnd(int data) {
        Node newNode = new Node(data);
        if this.end != null {
            this.end.left = newNode;
            newNode.right = this.end;
            this.end = newNode;
        } else {
            // the case of head not null but end as null not exist; so only the case of head and end as null
            if this.head != null {
                this.head.right = newNode;
                newNode.left = this.head;
                this.end = newNode;
            }else {
                this.head = newNode
            }
        }
    }
    public removeAtHead() {
        if this.head != null {
            this.head = this.head.right;
            this.head.left = null;
        }else {
            system.out.println("the linked list has no items to remove")
        }
    }

    public remoteAtEnd() {
        if this.end != null {
            // this.head must exist
            this.end = null;
            Node temp = this.head
            while (temp.right != null) {
                temp = temp.right
            }
            this.end == temp;
        }
    }

    public removeData(int data) {
        if this.head == null {
            return;
        }
        if this.head.data == data {
            this.head = this.head.right;
            this.head.left = null
            return;
        }
        Node temp = this.head
        while (temp.right != null) and temp.right.data != data  {
            temp = temp.right
            
        }
        if temp.right.right != null {
            temp.right = temp.right.right
            temp.right.left = temp
        }
        return;
    }
}