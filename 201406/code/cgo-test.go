package main

/*
#include <stdio.h>
#include <stdint.h>
#include <stdlib.h>

#define SJ_BUCKET_SIZE   5  // The maximum number of terms that the system can hold
#define SJ_NUM_BUCKETS   10

struct sjrev {
	uint32_t documentId;   // The document the term is in
	uint16_t rank;         // The rank of the term in the document
	struct sjrev *next;    // The next reverse index record
};

struct sjterm {
	struct sjrev *start;   // The first reverse index record
};

struct sjterm *buckets[SJ_NUM_BUCKETS] = { NULL };

struct sjterm *Term(uint32_t termId) {
	uint32_t bucket = termId/SJ_BUCKET_SIZE;
	if (bucket < SJ_NUM_BUCKETS) {
		if (buckets[bucket] == NULL) {
			// Initialise bucket
			buckets[bucket] = (struct sjterm *)malloc(SJ_BUCKET_SIZE*sizeof(struct sjterm));
		}
		return &buckets[bucket][termId-(bucket*SJ_BUCKET_SIZE)];
	} else {
		printf("Term(%d) = <exceeded_size>\n", termId); // Panic
		return NULL;
	}
}

// Add a reverse index record to a term
void AddRev(uint32_t termId, uint32_t documentId, uint16_t rank) {
	struct sjterm *term = Term(termId);
	if (term == NULL) return; // Couldn't initialise bucket
	struct sjrev *rev = (struct sjrev *)malloc(sizeof(struct sjrev));
	rev->documentId = documentId;
	rev->rank = rank;
	rev->next = term->start;
	term->start = rev;
}

// Remove a reverse index record for a term & document
void RemoveRev(uint32_t termId, uint16_t documentId) {
	struct sjterm *term = Term(termId);
	if (term == NULL) return; // Couldn't initialise bucket
	if (term->start == NULL) return; // No reverse index records
	struct sjrev *rev = term->start;
	if (rev->documentId == documentId) {
		term->start = rev->next;
		free(rev);
	} else {
		while (rev->next != NULL) {
			if (rev->next->documentId == documentId) {
				struct sjrev *tmp = rev->next;
				rev->next = rev->next->next;
				free(tmp);
				return;
			}
			rev = rev->next;
		}
	}
}

// Retreive the reverse index records for a term
void GetRev(uint32_t termId, uint16_t maxrank) {
	struct sjterm *term = Term(termId);
	if (term == NULL) return; // Couldn't initialise bucket
	printf("GetRev(%d)\n", termId);
	struct sjrev *rev = term->start;
	printf("Term(%d)->[", termId);
	while (rev != NULL) {
		if (rev->rank <= maxrank) {
			printf(" %d,%d", rev->documentId, rev->rank);
		}
		rev = rev->next;
	}
	printf(" ]\n");
}

// Dump out the storage (DEBUG)
void Dump() {
	struct sjrev *rev;
	int bucket, term;
	for (bucket = 0; bucket < SJ_NUM_BUCKETS; bucket++) {
		if (buckets[bucket] != NULL) {
			printf("Bucket[%d]:\n", bucket);
			for (term = 0; term < SJ_BUCKET_SIZE; term++) {
				printf("  Term[%d]:", (bucket*SJ_BUCKET_SIZE)+term);
				rev = buckets[bucket][term].start;
				while (rev != NULL) {
					printf(" %d(%d)", rev->documentId, rev->rank);
					rev = rev->next;
				}
				printf("\n");
			}
		}
	}
}

*/
import "C"

func main() {
	C.AddRev(4, 12, 1);
	C.AddRev(1, 12, 1);
	C.AddRev(2, 12, 1);
	C.AddRev(3, 12, 1);
	C.AddRev(4, 12, 1);
	C.AddRev(5, 12, 1);
	C.AddRev(6, 12, 1);
	C.AddRev(7, 12, 1);
	C.AddRev(4, 13, 2);
	C.GetRev(4,1);
	C.RemoveRev(4, 12);
	C.GetRev(4,3);
	C.AddRev(4, 14, 3);
	C.GetRev(4,5);
	C.RemoveRev(4, 14);
	C.GetRev(4,5);
	C.Dump();
}

