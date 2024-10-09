class ListNode {
	int val;
	ListNode next;

	public ListNode(int val) {
		this.val = val;
		this.next = null;
	}
}

public class Queue {
	ListNode left; // front of Queue front -> [1,2,3]
	ListNode right; // back of Queue

	public Queue() {
		this.left = null;
		this.right = null;
	}

	// add to front of queue
	public void enqueue(int value) {
		ListNode newNode = new ListNode(value);
		if (this.right != null) {
			// Queue is not empty
			this.right.next = newNode;
			this.right = this.right.next;
		} else {
			// Queue is empty
			this.left = newNode;
			this.right = newNode;
		}
	}

	// remove from back of queue
	public int dequeue() {
		if (this.left == null) {
			// Queue is empty
			System.exit(0);
		}

		int value = this.left.val;
		this.left = this.left.next;
		if (this.left == null) {
			this.right = null;
		}

		return value;

	}

	public void print() {
		ListNode current = this.left;
		while (current != null) {
			System.out.print(current.val + " -> ");
			current = current.next;
		}
		System.out.println();

	}

	public static void main(String[] args) {
		Queue q = new Queue();
		q.enqueue(1);
		q.enqueue(2);
		q.enqueue(3);

		q.print();

		q.dequeue();
		q.print();
	}

}
