@0xb3a837315d527939;
using Go = import "go.capnp";
$Go.package("main");
$Go.import("testpkg");


struct ClueCapn { 
   term     @0:   Text; 
   intro    @1:   Text; 
   potency  @2:   Float32; 
} 

struct RevCapn { 
   docId   @0:   UInt32; 
   rank    @1:   UInt16; 
   inMeta  @2:   UInt16; 
} 

struct RevsCapn { 
   data  @0:   List(RevCapn); 
} 

struct ShotgunCapn { 
   term     @0:   Text; 
   potency  @1:   Float32; 
} 

struct TermCapn { 
   termStr          @0:   Text; 
   slot             @1:   UInt32; 
   numDocuments     @2:   UInt32; 
   numWords         @3:   UInt8; 
   shotgun          @4:   List(ShotgunCapn); 
   clues            @5:   List(ClueCapn); 
   interactionsPos  @6:   UInt16; 
   interactionsNeg  @7:   UInt16; 
   hardcodedScore   @8:   Int8; 
   infogain         @9:   Float32; 
} 

##compile with:

##
##
##   capnp compile -ogo ./schema.capnp

