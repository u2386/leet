import java.util.HashMap;
import java.util.Map;

/**
 * Your LRUCache object will be instantiated and called as such: LRUCache obj = new
 * LRUCache(capacity); int param_1 = obj.get(key); obj.put(key,value);
 */
class LRUCache {

  private int capacity;
  private Map<Integer, Node> kv;
  private Node head;
  private Node tail;

  public LRUCache(int capacity) {
    this.capacity = capacity;
    this.kv = new HashMap<>();
    this.head = new Node();
    this.tail = new Node();
    this.head.next = this.tail;
    this.tail.prev = this.head;
  }

  public int get(int key) {
    Node n = kv.get(key);
    if (n != null) {
      remove(n);
      add(n);
      return n.value;
    }
    return -1;
  }

  public void put(int key, int value) {
    Node n;
    if ((n = kv.get(key)) != null) {
      remove(n);
      add(n);
      n.value = value;
      kv.put(key, n);
    } else {
      if (kv.size() >= capacity) {
        removeFromKV(tail.prev.key);
        remove(tail.prev);
      }
      n = new Node(key, value);
      add(n);
      kv.put(key, n);
    }
  }

  private void add(Node n) {
    n.prev = head;
    n.next = head.next;
    head.next = n;
    n.next.prev = n;
  }

  private void remove(Node n) {
    n.prev.next = n.next;
    n.next.prev = n.prev;
    n.prev = n.next = null;
  }

  private void removeFromKV(int key) {
    kv.put(key, null);
  }

  private static class Node {

    private Node prev;
    private Node next;
    private int key;
    private int value;

    Node() {
      this.prev = this.next = null;
    }

    Node(int key, int value) {
      this.key = key;
      this.value = value;
      this.prev = this.next = null;
    }
  }
}
