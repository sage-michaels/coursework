#include <stdio.h>
#include <stdlib.h>
#include "Cache.h"


/* This method configures the cache.  The capacity and block size need
 * to be specified in bytes.  All of the three configuration
 * arguments must be powers of two.  Additionally, the
 * capacity divided by the (block_size * associativity) must not
 * have a fractional component.  It should initialize the statistics
 * to zero.  Finally, the method prints out
 * information about the configuration of the cache (see example
 * output files).
 */


void initializeCache(Cache *c, int capacity, int block_size, int associativity){
  c -> num_sets = capacity/(block_size * associativity);
  c -> assoc = associativity; 
  int num_sets = c->num_sets;
  c -> sets = malloc(num_sets*sizeof(Set));
  Set *sets = c -> sets;
  int assoc = c -> assoc;
  int i;
  for(i = 0;i <c-> num_sets;i++){
    sets[i].blocks = malloc((assoc)*sizeof(Block));
    sets[i].index = i;
    int j;
    Block* block = sets[i].blocks;
    for(j=0;j<(assoc);j++){
      block[j].lru = -1;
      block[j].tag = 0;
    } 
  }
  c->accessStatistics.accesses = 0;
  c->accessStatistics.hits     = 0;
  c->accessStatistics.misses   = 0;
}
 


/* 
 * This function releases all memory dynamically allocated by
 * initializeCache()
 */
void destroyCache(Cache *c){
  int i;
  for( i = 0; i<(c->num_sets);i++){
    Block *block =(c->sets[i].blocks);
    free(block);
  }  
  free(c->sets);
}


/* Adds the address specified as an access to the current cache contents, 
 * potentially evicting an existing entry in the cache.  The address
 * is specified in bytes.
 */
void addAccess(Cache *c, long long addr){
  c->accessStatistics.accesses++;
  long long set_addr = addr / c->num_sets;
  long long index    = set_addr %(c-> num_sets);
  int tag           = set_addr / (c->num_sets);
  Block *location     = c ->sets[index].blocks;
  int i             =  0;
  bool hit         = false;
  bool entered     = false;
  int evict_val     = -1;
  int evict_loc     = -1;
  while(i<(c->assoc)){
    Block *block = &(location[i]);
    if(evict_val<block->lru){
      evict_loc = i;
      evict_val = block->lru;
    }
    if(entered){
      if(block->lru>-1){
        block->lru +=1 ;
      }
      i++;
    }else{
      if(block->tag ==tag){
        if(block->lru >-1){
          hit = true;
        }else{
          block->lru = 0;
          block->tag = tag;
          entered = true;
        }
      }else{
        if(block->lru>-1){
          block->lru += 1;
        }else{
          block->lru = 0;
          block->tag = tag;
          entered = true;
        }
      }
      i++;  
    }
  }    
  if(hit){
    c->accessStatistics.hits++;
  }else{
    if(!entered){
      location[evict_loc].tag = tag;
      location[evict_loc].lru = 0;
    }
    c->accessStatistics.misses++;
  }
}
  
  
/* Prints the current contents of the cache.  The output is organized
 * by cache sets and displays the index, tag, valid bit, and lru state.
 */
void printContents(Cache *c){
  int i;
  for(i = 0;i<(c->num_sets);i++){
    printf("****** SET %i****************************\n", i);
    Set location = c->sets[i];
    int j;
    for(j = 0; j<(c->assoc); j++){
      int lru = location.blocks[j].lru;
      int valid;
      if(lru == -1){
        valid = 0;
      }else{
        valid = 1;
      }
      int tag = location.blocks[j].tag;
      int index;
      if(c->assoc == 0){
        index = i;
      }else{
        index = j;
      }
      printf("Index %d: tag 0x%x valid %i lru %i\n",index,tag,valid,lru);
    }
    printf("****************************************\n");
  } 
}

/* Prints the statistics about cache accesses including the number of
 * accesses, hits, misses, and hit rate.
 */
void printStatistics(Cache *c){
  double hits =(double)  c->accessStatistics.hits;
  printf("ACCESSES %lld \n",getNumberAccesses(c));
  printf("HITS %i \n",(int)  hits);
  printf("MISSES %lld \n",getNumberMisses(c));
  printf("HIT RATE %lf\n", (double) hits/getNumberAccesses(c));
}

/* 
 * Returns the number of accesses made to the cache. 
 */
long long getNumberAccesses(Cache *c){
  long long n = c->accessStatistics.accesses;
  return n;
}

/* 
 * Returns the number of access that resulted in cache misses.
 */
long long getNumberMisses(Cache *c){
  long long n = c->accessStatistics.misses;
  return n;
}