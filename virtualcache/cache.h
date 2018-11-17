#ifndef _CACHE_H_
#define _CACHE_H_

#include <stdio.h>
#include <stdlib.h>
#include <math.h>

/*
 * This struct represents a configurable memory cache.  To configure the
 * cache, the user must specify the cache capacity in bytes, the block
 * size in bytes, and the set-associativity.  
 *
 * The user can add memory accesses to the cache by specifying the byte
 * address for the memory location.  Each access is assumed to be 4 
 * bytes wide.  
 *
 * The user has the ability to print out the current contents of the cache
 * as well as query the cache about access and hit/miss statistics.
 */

/*
 * This struct is used to keep information about hits and misses.
 */
typedef struct statistics {
  long long accesses;
  long long hits;
  long long misses;
} Statistics;


/* 
 * Created to make it possible to use bool as a type with false=0 and true=1
 */
typedef int bool;
enum BOOL{false, true};

/* 
 * This struct represents a single block in the cache 
 */
typedef struct block {
  int lru;
  int tag;
}Block;

/* 
 * This struct represents a set in a cache.  For direct-mapped caches,
 * there is one block per set.  For fully-associative caches, all
 * of the blocks in the cache are in a single set.
 */
typedef struct set {
 Block *blocks; 
 int index;
}Set;

/* 
 * This struct represents a cache.  In particular, a cache contains an
 * array of Sets. 
 */
typedef struct cache {
  Set *sets;    // This must be your primary data structure.  
  int num_sets;
  int assoc;
  // configuration (e.g. capacity, set associativity, etc.) and potentially
  // masks you create for pulling fields out of memory addresses.

  Statistics accessStatistics;
}Cache;


/* This method configures the cache.  The capacity and block size need
 * to be specified in bytes.  All of the three configuration
 * arguments must be powers of two.  Additionally, the
 * capacity divided by the (block_size * associativity) must not
 * have a fractional component.  It should initialize the statistics
 * to zero.  Finally, the method prints out
 * information about the configuration of the cache (see example
 * output files).
 */
void initializeCache(Cache *c, int capacity, int block_size, int associativity);  

/* 
 * This function releases all memory dynamically allocated by
 * initializeCache()
 */
void destroyCache(Cache *c);


/* Adds the address specified as an access to the current cache contents, 
 * potentially evicting an existing entry in the cache.  The address
 * is specified in bytes.
 */
void addAccess(Cache *c, long long addr);
  
/* Prints the current contents of the cache.  The output is organized
 * by cache sets and displays the index, tag, valid bit, and lru state.
 */
void printContents(Cache *c);

/* Prints the statistics about cache accesses including the number of
 * accesses, hits, misses, and hit rate.
 */
void printStatistics(Cache *c);

/* 
 * Returns the number of accesses made to the cache. 
 */
long long getNumberAccesses(Cache *c);

/* 
 * Returns the number of access that resulted in cache misses.
 */
long long getNumberMisses(Cache *c);

#endif