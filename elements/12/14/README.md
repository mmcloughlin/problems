Find a missing IP address.

Given a file of approx 1 billion IP addresses as 32-bit integers. Find one
that's not in the file. Suppose you have unlimited drive space but just a few
megabytes of RAM.

---

1 billion IP addresses is approximately 2^30 4 byte quantities. That's 4 GiB.

* Bucket by first two bytes, look for one with less than 2^24 IPs in it.
* Within that, look for a missing IP.
