package threefish

import (
    "os"
    "encoding/binary"
)

// Constants for Threefish 256 implementation
//
const (
    CIPHER_SIZE_256 = 256
    CIPHER_QWORDS_256 = CIPHER_SIZE_256 / 64
    EXPANDED_KEY_SIZE_256 = CIPHER_QWORDS_256 + 1
)

type threefish256 struct {
    expanedTweak [EXPANDED_TWEAK_SIZE]uint64
    expanedKey [EXPANDED_KEY_SIZE_256]uint64
    tmpData1, tmpData2  [CIPHER_QWORDS_256]uint64
}

// Get an initialized Threefish256 structure
func newThreefish256(key []byte, tweak []uint64) (*threefish256, os.Error) {
    c := new(threefish256)

    // Initialize tweak 
    setTweak(tweak, c.expanedTweak[:])

    c.expanedKey[EXPANDED_KEY_SIZE_256 - 1] = KEY_SCHEDULE_CONST

    // Convert key bytes to required uint64 array
    if key != nil {
        var tmpkey [EXPANDED_KEY_SIZE_256]uint64        
        for i := 0; i < EXPANDED_KEY_SIZE_256 - 1; i++ {
            tmpkey[i] = binary.LittleEndian.Uint64(key[i*8:i*8+8])
        }
        setKey(tmpkey[:], c.expanedKey[:])
    }
    return c, nil
}

// Get an initialized Threefish256 structure
func newThreefish256_64(key, tweak []uint64) (*threefish256, os.Error) {
    c := new(threefish256)

    // Initialize tweak 
    setTweak(tweak, c.expanedTweak[:])

    c.expanedKey[EXPANDED_KEY_SIZE_256 - 1] = KEY_SCHEDULE_CONST

    setKey(key, c.expanedKey[:])
    return c, nil
}

func (tf *threefish256) getTempData() ([]uint64, []uint64) {
    return tf.tmpData1[:], tf.tmpData2[:]
}

func (tf *threefish256) setTweak(tweak []uint64) {
    setTweak(tweak, tf.expanedTweak[:])    
}

func (tf *threefish256) setKey(key []uint64) {
    setKey(key, tf.expanedKey[:])    
}

func (tf *threefish256) encrypt(input, output [] uint64) {
    
   b0 := input[0]
   b1 := input[1]
   b2 := input[2]
   b3 := input[3]
   
   k0 := tf.expanedKey[0]
   k1 := tf.expanedKey[1]
   k2 := tf.expanedKey[2] 
   k3 := tf.expanedKey[3]
   k4 := tf.expanedKey[4]
   
   t0 := tf.expanedTweak[0]
   t1 := tf.expanedTweak[1]
   t2 := tf.expanedTweak[2]
   
   b1 += k1 + t0; b0 += b1 + k0; b1 = ((b1 << 14) | (b1 >> (64 - 14))) ^ b0;
   b3 += k3; b2 += b3 + k2 + t1; b3 = ((b3 << 16) | (b3 >> (64 - 16))) ^ b2;
   b0 += b3; b3 = ((b3 << 52) | (b3 >> (64 - 52))) ^ b0;
   b2 += b1; b1 = ((b1 << 57) | (b1 >> (64 - 57))) ^ b2;
   b0 += b1; b1 = ((b1 << 23) | (b1 >> (64 - 23))) ^ b0;
   b2 += b3; b3 = ((b3 << 40) | (b3 >> (64 - 40))) ^ b2;
   b0 += b3; b3 = ((b3 << 5) | (b3 >> (64 - 5))) ^ b0;
   b2 += b1; b1 = ((b1 << 37) | (b1 >> (64 - 37))) ^ b2;
   b1 += k2 + t1; b0 += b1 + k1; b1 = ((b1 << 25) | (b1 >> (64 - 25))) ^ b0;
   b3 += k4 + 1; b2 += b3 + k3 + t2; b3 = ((b3 << 33) | (b3 >> (64 - 33))) ^ b2;
   b0 += b3; b3 = ((b3 << 46) | (b3 >> (64 - 46))) ^ b0;
   b2 += b1; b1 = ((b1 << 12) | (b1 >> (64 - 12))) ^ b2;
   b0 += b1; b1 = ((b1 << 58) | (b1 >> (64 - 58))) ^ b0;
   b2 += b3; b3 = ((b3 << 22) | (b3 >> (64 - 22))) ^ b2;
   b0 += b3; b3 = ((b3 << 32) | (b3 >> (64 - 32))) ^ b0;
   b2 += b1; b1 = ((b1 << 32) | (b1 >> (64 - 32))) ^ b2;
   b1 += k3 + t2; b0 += b1 + k2; b1 = ((b1 << 14) | (b1 >> (64 - 14))) ^ b0;
   b3 += k0 + 2; b2 += b3 + k4 + t0; b3 = ((b3 << 16) | (b3 >> (64 - 16))) ^ b2;
   b0 += b3; b3 = ((b3 << 52) | (b3 >> (64 - 52))) ^ b0;
   b2 += b1; b1 = ((b1 << 57) | (b1 >> (64 - 57))) ^ b2;
   b0 += b1; b1 = ((b1 << 23) | (b1 >> (64 - 23))) ^ b0;
   b2 += b3; b3 = ((b3 << 40) | (b3 >> (64 - 40))) ^ b2;
   b0 += b3; b3 = ((b3 << 5) | (b3 >> (64 - 5))) ^ b0;
   b2 += b1; b1 = ((b1 << 37) | (b1 >> (64 - 37))) ^ b2;
   b1 += k4 + t0; b0 += b1 + k3; b1 = ((b1 << 25) | (b1 >> (64 - 25))) ^ b0;
   b3 += k1 + 3; b2 += b3 + k0 + t1; b3 = ((b3 << 33) | (b3 >> (64 - 33))) ^ b2;
   b0 += b3; b3 = ((b3 << 46) | (b3 >> (64 - 46))) ^ b0;
   b2 += b1; b1 = ((b1 << 12) | (b1 >> (64 - 12))) ^ b2;
   b0 += b1; b1 = ((b1 << 58) | (b1 >> (64 - 58))) ^ b0;
   b2 += b3; b3 = ((b3 << 22) | (b3 >> (64 - 22))) ^ b2;
   b0 += b3; b3 = ((b3 << 32) | (b3 >> (64 - 32))) ^ b0;
   b2 += b1; b1 = ((b1 << 32) | (b1 >> (64 - 32))) ^ b2;
   b1 += k0 + t1; b0 += b1 + k4; b1 = ((b1 << 14) | (b1 >> (64 - 14))) ^ b0;
   b3 += k2 + 4; b2 += b3 + k1 + t2; b3 = ((b3 << 16) | (b3 >> (64 - 16))) ^ b2;
   b0 += b3; b3 = ((b3 << 52) | (b3 >> (64 - 52))) ^ b0;
   b2 += b1; b1 = ((b1 << 57) | (b1 >> (64 - 57))) ^ b2;
   b0 += b1; b1 = ((b1 << 23) | (b1 >> (64 - 23))) ^ b0;
   b2 += b3; b3 = ((b3 << 40) | (b3 >> (64 - 40))) ^ b2;
   b0 += b3; b3 = ((b3 << 5) | (b3 >> (64 - 5))) ^ b0;
   b2 += b1; b1 = ((b1 << 37) | (b1 >> (64 - 37))) ^ b2;
   b1 += k1 + t2; b0 += b1 + k0; b1 = ((b1 << 25) | (b1 >> (64 - 25))) ^ b0;
   b3 += k3 + 5; b2 += b3 + k2 + t0; b3 = ((b3 << 33) | (b3 >> (64 - 33))) ^ b2;
   b0 += b3; b3 = ((b3 << 46) | (b3 >> (64 - 46))) ^ b0;
   b2 += b1; b1 = ((b1 << 12) | (b1 >> (64 - 12))) ^ b2;
   b0 += b1; b1 = ((b1 << 58) | (b1 >> (64 - 58))) ^ b0;
   b2 += b3; b3 = ((b3 << 22) | (b3 >> (64 - 22))) ^ b2;
   b0 += b3; b3 = ((b3 << 32) | (b3 >> (64 - 32))) ^ b0;
   b2 += b1; b1 = ((b1 << 32) | (b1 >> (64 - 32))) ^ b2;
   b1 += k2 + t0; b0 += b1 + k1; b1 = ((b1 << 14) | (b1 >> (64 - 14))) ^ b0;
   b3 += k4 + 6; b2 += b3 + k3 + t1; b3 = ((b3 << 16) | (b3 >> (64 - 16))) ^ b2;
   b0 += b3; b3 = ((b3 << 52) | (b3 >> (64 - 52))) ^ b0;
   b2 += b1; b1 = ((b1 << 57) | (b1 >> (64 - 57))) ^ b2;
   b0 += b1; b1 = ((b1 << 23) | (b1 >> (64 - 23))) ^ b0;
   b2 += b3; b3 = ((b3 << 40) | (b3 >> (64 - 40))) ^ b2;
   b0 += b3; b3 = ((b3 << 5) | (b3 >> (64 - 5))) ^ b0;
   b2 += b1; b1 = ((b1 << 37) | (b1 >> (64 - 37))) ^ b2;
   b1 += k3 + t1; b0 += b1 + k2; b1 = ((b1 << 25) | (b1 >> (64 - 25))) ^ b0;
   b3 += k0 + 7; b2 += b3 + k4 + t2; b3 = ((b3 << 33) | (b3 >> (64 - 33))) ^ b2;
   b0 += b3; b3 = ((b3 << 46) | (b3 >> (64 - 46))) ^ b0;
   b2 += b1; b1 = ((b1 << 12) | (b1 >> (64 - 12))) ^ b2;
   b0 += b1; b1 = ((b1 << 58) | (b1 >> (64 - 58))) ^ b0;
   b2 += b3; b3 = ((b3 << 22) | (b3 >> (64 - 22))) ^ b2;
   b0 += b3; b3 = ((b3 << 32) | (b3 >> (64 - 32))) ^ b0;
   b2 += b1; b1 = ((b1 << 32) | (b1 >> (64 - 32))) ^ b2;
   b1 += k4 + t2; b0 += b1 + k3; b1 = ((b1 << 14) | (b1 >> (64 - 14))) ^ b0;
   b3 += k1 + 8; b2 += b3 + k0 + t0; b3 = ((b3 << 16) | (b3 >> (64 - 16))) ^ b2;
   b0 += b3; b3 = ((b3 << 52) | (b3 >> (64 - 52))) ^ b0;
   b2 += b1; b1 = ((b1 << 57) | (b1 >> (64 - 57))) ^ b2;
   b0 += b1; b1 = ((b1 << 23) | (b1 >> (64 - 23))) ^ b0;
   b2 += b3; b3 = ((b3 << 40) | (b3 >> (64 - 40))) ^ b2;
   b0 += b3; b3 = ((b3 << 5) | (b3 >> (64 - 5))) ^ b0;
   b2 += b1; b1 = ((b1 << 37) | (b1 >> (64 - 37))) ^ b2;
   b1 += k0 + t0; b0 += b1 + k4; b1 = ((b1 << 25) | (b1 >> (64 - 25))) ^ b0;
   b3 += k2 + 9; b2 += b3 + k1 + t1; b3 = ((b3 << 33) | (b3 >> (64 - 33))) ^ b2;
   b0 += b3; b3 = ((b3 << 46) | (b3 >> (64 - 46))) ^ b0;
   b2 += b1; b1 = ((b1 << 12) | (b1 >> (64 - 12))) ^ b2;
   b0 += b1; b1 = ((b1 << 58) | (b1 >> (64 - 58))) ^ b0;
   b2 += b3; b3 = ((b3 << 22) | (b3 >> (64 - 22))) ^ b2;
   b0 += b3; b3 = ((b3 << 32) | (b3 >> (64 - 32))) ^ b0;
   b2 += b1; b1 = ((b1 << 32) | (b1 >> (64 - 32))) ^ b2;
   b1 += k1 + t1; b0 += b1 + k0; b1 = ((b1 << 14) | (b1 >> (64 - 14))) ^ b0;
   b3 += k3 + 10; b2 += b3 + k2 + t2; b3 = ((b3 << 16) | (b3 >> (64 - 16))) ^ b2;
   b0 += b3; b3 = ((b3 << 52) | (b3 >> (64 - 52))) ^ b0;
   b2 += b1; b1 = ((b1 << 57) | (b1 >> (64 - 57))) ^ b2;
   b0 += b1; b1 = ((b1 << 23) | (b1 >> (64 - 23))) ^ b0;
   b2 += b3; b3 = ((b3 << 40) | (b3 >> (64 - 40))) ^ b2;
   b0 += b3; b3 = ((b3 << 5) | (b3 >> (64 - 5))) ^ b0;
   b2 += b1; b1 = ((b1 << 37) | (b1 >> (64 - 37))) ^ b2;
   b1 += k2 + t2; b0 += b1 + k1; b1 = ((b1 << 25) | (b1 >> (64 - 25))) ^ b0;
   b3 += k4 + 11; b2 += b3 + k3 + t0; b3 = ((b3 << 33) | (b3 >> (64 - 33))) ^ b2;
   b0 += b3; b3 = ((b3 << 46) | (b3 >> (64 - 46))) ^ b0;
   b2 += b1; b1 = ((b1 << 12) | (b1 >> (64 - 12))) ^ b2;
   b0 += b1; b1 = ((b1 << 58) | (b1 >> (64 - 58))) ^ b0;
   b2 += b3; b3 = ((b3 << 22) | (b3 >> (64 - 22))) ^ b2;
   b0 += b3; b3 = ((b3 << 32) | (b3 >> (64 - 32))) ^ b0;
   b2 += b1; b1 = ((b1 << 32) | (b1 >> (64 - 32))) ^ b2;
   b1 += k3 + t0; b0 += b1 + k2; b1 = ((b1 << 14) | (b1 >> (64 - 14))) ^ b0;
   b3 += k0 + 12; b2 += b3 + k4 + t1; b3 = ((b3 << 16) | (b3 >> (64 - 16))) ^ b2;
   b0 += b3; b3 = ((b3 << 52) | (b3 >> (64 - 52))) ^ b0;
   b2 += b1; b1 = ((b1 << 57) | (b1 >> (64 - 57))) ^ b2;
   b0 += b1; b1 = ((b1 << 23) | (b1 >> (64 - 23))) ^ b0;
   b2 += b3; b3 = ((b3 << 40) | (b3 >> (64 - 40))) ^ b2;
   b0 += b3; b3 = ((b3 << 5) | (b3 >> (64 - 5))) ^ b0;
   b2 += b1; b1 = ((b1 << 37) | (b1 >> (64 - 37))) ^ b2;
   b1 += k4 + t1; b0 += b1 + k3; b1 = ((b1 << 25) | (b1 >> (64 - 25))) ^ b0;
   b3 += k1 + 13; b2 += b3 + k0 + t2; b3 = ((b3 << 33) | (b3 >> (64 - 33))) ^ b2;
   b0 += b3; b3 = ((b3 << 46) | (b3 >> (64 - 46))) ^ b0;
   b2 += b1; b1 = ((b1 << 12) | (b1 >> (64 - 12))) ^ b2;
   b0 += b1; b1 = ((b1 << 58) | (b1 >> (64 - 58))) ^ b0;
   b2 += b3; b3 = ((b3 << 22) | (b3 >> (64 - 22))) ^ b2;
   b0 += b3; b3 = ((b3 << 32) | (b3 >> (64 - 32))) ^ b0;
   b2 += b1; b1 = ((b1 << 32) | (b1 >> (64 - 32))) ^ b2;
   b1 += k0 + t2; b0 += b1 + k4; b1 = ((b1 << 14) | (b1 >> (64 - 14))) ^ b0;
   b3 += k2 + 14; b2 += b3 + k1 + t0; b3 = ((b3 << 16) | (b3 >> (64 - 16))) ^ b2;
   b0 += b3; b3 = ((b3 << 52) | (b3 >> (64 - 52))) ^ b0;
   b2 += b1; b1 = ((b1 << 57) | (b1 >> (64 - 57))) ^ b2;
   b0 += b1; b1 = ((b1 << 23) | (b1 >> (64 - 23))) ^ b0;
   b2 += b3; b3 = ((b3 << 40) | (b3 >> (64 - 40))) ^ b2;
   b0 += b3; b3 = ((b3 << 5) | (b3 >> (64 - 5))) ^ b0;
   b2 += b1; b1 = ((b1 << 37) | (b1 >> (64 - 37))) ^ b2;
   b1 += k1 + t0; b0 += b1 + k0; b1 = ((b1 << 25) | (b1 >> (64 - 25))) ^ b0;
   b3 += k3 + 15; b2 += b3 + k2 + t1; b3 = ((b3 << 33) | (b3 >> (64 - 33))) ^ b2;
   b0 += b3; b3 = ((b3 << 46) | (b3 >> (64 - 46))) ^ b0;
   b2 += b1; b1 = ((b1 << 12) | (b1 >> (64 - 12))) ^ b2;
   b0 += b1; b1 = ((b1 << 58) | (b1 >> (64 - 58))) ^ b0;
   b2 += b3; b3 = ((b3 << 22) | (b3 >> (64 - 22))) ^ b2;
   b0 += b3; b3 = ((b3 << 32) | (b3 >> (64 - 32))) ^ b0;
   b2 += b1; b1 = ((b1 << 32) | (b1 >> (64 - 32))) ^ b2;
   b1 += k2 + t1; b0 += b1 + k1; b1 = ((b1 << 14) | (b1 >> (64 - 14))) ^ b0;
   b3 += k4 + 16; b2 += b3 + k3 + t2; b3 = ((b3 << 16) | (b3 >> (64 - 16))) ^ b2;
   b0 += b3; b3 = ((b3 << 52) | (b3 >> (64 - 52))) ^ b0;
   b2 += b1; b1 = ((b1 << 57) | (b1 >> (64 - 57))) ^ b2;
   b0 += b1; b1 = ((b1 << 23) | (b1 >> (64 - 23))) ^ b0;
   b2 += b3; b3 = ((b3 << 40) | (b3 >> (64 - 40))) ^ b2;
   b0 += b3; b3 = ((b3 << 5) | (b3 >> (64 - 5))) ^ b0;
   b2 += b1; b1 = ((b1 << 37) | (b1 >> (64 - 37))) ^ b2;
   b1 += k3 + t2; b0 += b1 + k2; b1 = ((b1 << 25) | (b1 >> (64 - 25))) ^ b0;
   b3 += k0 + 17; b2 += b3 + k4 + t0; b3 = ((b3 << 33) | (b3 >> (64 - 33))) ^ b2;
   b0 += b3; b3 = ((b3 << 46) | (b3 >> (64 - 46))) ^ b0;
   b2 += b1; b1 = ((b1 << 12) | (b1 >> (64 - 12))) ^ b2;
   b0 += b1; b1 = ((b1 << 58) | (b1 >> (64 - 58))) ^ b0;
   b2 += b3; b3 = ((b3 << 22) | (b3 >> (64 - 22))) ^ b2;
   b0 += b3; b3 = ((b3 << 32) | (b3 >> (64 - 32))) ^ b0;
   b2 += b1; b1 = ((b1 << 32) | (b1 >> (64 - 32))) ^ b2;

   output[0] = b0 + k3;
   output[1] = b1 + k4 + t0;
   output[2] = b2 + k0 + t1;
   output[3] = b3 + k1 + 18;

}

func (tf *threefish256) decrypt(input, output []uint64) {
    
   b0 := input[0]
   b1 := input[1]
   b2 := input[2]
   b3 := input[3]
   
   k0 := tf.expanedKey[0]
   k1 := tf.expanedKey[1]
   k2 := tf.expanedKey[2] 
   k3 := tf.expanedKey[3]
   k4 := tf.expanedKey[4]
   
   t0 := tf.expanedTweak[0]
   t1 := tf.expanedTweak[1]
   t2 := tf.expanedTweak[2]

   var tmp uint64

   b0 -= k3;
   b1 -= k4 + t0;
   b2 -= k0 + t1;
   b3 -= k1 + 18;
   tmp = b3 ^ b0; b3 = (tmp >> 32) | (tmp << (64 - 32)); b0 -= b3;
   tmp = b1 ^ b2; b1 = (tmp >> 32) | (tmp << (64 - 32)); b2 -= b1;
   tmp = b1 ^ b0; b1 = (tmp >> 58) | (tmp << (64 - 58)); b0 -= b1;
   tmp = b3 ^ b2; b3 = (tmp >> 22) | (tmp << (64 - 22)); b2 -= b3;
   tmp = b3 ^ b0; b3 = (tmp >> 46) | (tmp << (64 - 46)); b0 -= b3;
   tmp = b1 ^ b2; b1 = (tmp >> 12) | (tmp << (64 - 12)); b2 -= b1;
   tmp = b1 ^ b0; b1 = (tmp >> 25) | (tmp << (64 - 25)); b0 -= b1 + k2; b1 -= k3 + t2;
   tmp = b3 ^ b2; b3 = (tmp >> 33) | (tmp << (64 - 33)); b2 -= b3 + k4 + t0; b3 -= k0 + 17;
   tmp = b3 ^ b0; b3 = (tmp >> 5) | (tmp << (64 - 5)); b0 -= b3;
   tmp = b1 ^ b2; b1 = (tmp >> 37) | (tmp << (64 - 37)); b2 -= b1;
   tmp = b1 ^ b0; b1 = (tmp >> 23) | (tmp << (64 - 23)); b0 -= b1;
   tmp = b3 ^ b2; b3 = (tmp >> 40) | (tmp << (64 - 40)); b2 -= b3;
   tmp = b3 ^ b0; b3 = (tmp >> 52) | (tmp << (64 - 52)); b0 -= b3;
   tmp = b1 ^ b2; b1 = (tmp >> 57) | (tmp << (64 - 57)); b2 -= b1;
   tmp = b1 ^ b0; b1 = (tmp >> 14) | (tmp << (64 - 14)); b0 -= b1 + k1; b1 -= k2 + t1;
   tmp = b3 ^ b2; b3 = (tmp >> 16) | (tmp << (64 - 16)); b2 -= b3 + k3 + t2; b3 -= k4 + 16;
   tmp = b3 ^ b0; b3 = (tmp >> 32) | (tmp << (64 - 32)); b0 -= b3;
   tmp = b1 ^ b2; b1 = (tmp >> 32) | (tmp << (64 - 32)); b2 -= b1;
   tmp = b1 ^ b0; b1 = (tmp >> 58) | (tmp << (64 - 58)); b0 -= b1;
   tmp = b3 ^ b2; b3 = (tmp >> 22) | (tmp << (64 - 22)); b2 -= b3;
   tmp = b3 ^ b0; b3 = (tmp >> 46) | (tmp << (64 - 46)); b0 -= b3;
   tmp = b1 ^ b2; b1 = (tmp >> 12) | (tmp << (64 - 12)); b2 -= b1;
   tmp = b1 ^ b0; b1 = (tmp >> 25) | (tmp << (64 - 25)); b0 -= b1 + k0; b1 -= k1 + t0;
   tmp = b3 ^ b2; b3 = (tmp >> 33) | (tmp << (64 - 33)); b2 -= b3 + k2 + t1; b3 -= k3 + 15;
   tmp = b3 ^ b0; b3 = (tmp >> 5) | (tmp << (64 - 5)); b0 -= b3;
   tmp = b1 ^ b2; b1 = (tmp >> 37) | (tmp << (64 - 37)); b2 -= b1;
   tmp = b1 ^ b0; b1 = (tmp >> 23) | (tmp << (64 - 23)); b0 -= b1;
   tmp = b3 ^ b2; b3 = (tmp >> 40) | (tmp << (64 - 40)); b2 -= b3;
   tmp = b3 ^ b0; b3 = (tmp >> 52) | (tmp << (64 - 52)); b0 -= b3;
   tmp = b1 ^ b2; b1 = (tmp >> 57) | (tmp << (64 - 57)); b2 -= b1;
   tmp = b1 ^ b0; b1 = (tmp >> 14) | (tmp << (64 - 14)); b0 -= b1 + k4; b1 -= k0 + t2;
   tmp = b3 ^ b2; b3 = (tmp >> 16) | (tmp << (64 - 16)); b2 -= b3 + k1 + t0; b3 -= k2 + 14;
   tmp = b3 ^ b0; b3 = (tmp >> 32) | (tmp << (64 - 32)); b0 -= b3;
   tmp = b1 ^ b2; b1 = (tmp >> 32) | (tmp << (64 - 32)); b2 -= b1;
   tmp = b1 ^ b0; b1 = (tmp >> 58) | (tmp << (64 - 58)); b0 -= b1;
   tmp = b3 ^ b2; b3 = (tmp >> 22) | (tmp << (64 - 22)); b2 -= b3;
   tmp = b3 ^ b0; b3 = (tmp >> 46) | (tmp << (64 - 46)); b0 -= b3;
   tmp = b1 ^ b2; b1 = (tmp >> 12) | (tmp << (64 - 12)); b2 -= b1;
   tmp = b1 ^ b0; b1 = (tmp >> 25) | (tmp << (64 - 25)); b0 -= b1 + k3; b1 -= k4 + t1;
   tmp = b3 ^ b2; b3 = (tmp >> 33) | (tmp << (64 - 33)); b2 -= b3 + k0 + t2; b3 -= k1 + 13;
   tmp = b3 ^ b0; b3 = (tmp >> 5) | (tmp << (64 - 5)); b0 -= b3;
   tmp = b1 ^ b2; b1 = (tmp >> 37) | (tmp << (64 - 37)); b2 -= b1;
   tmp = b1 ^ b0; b1 = (tmp >> 23) | (tmp << (64 - 23)); b0 -= b1;
   tmp = b3 ^ b2; b3 = (tmp >> 40) | (tmp << (64 - 40)); b2 -= b3;
   tmp = b3 ^ b0; b3 = (tmp >> 52) | (tmp << (64 - 52)); b0 -= b3;
   tmp = b1 ^ b2; b1 = (tmp >> 57) | (tmp << (64 - 57)); b2 -= b1;
   tmp = b1 ^ b0; b1 = (tmp >> 14) | (tmp << (64 - 14)); b0 -= b1 + k2; b1 -= k3 + t0;
   tmp = b3 ^ b2; b3 = (tmp >> 16) | (tmp << (64 - 16)); b2 -= b3 + k4 + t1; b3 -= k0 + 12;
   tmp = b3 ^ b0; b3 = (tmp >> 32) | (tmp << (64 - 32)); b0 -= b3;
   tmp = b1 ^ b2; b1 = (tmp >> 32) | (tmp << (64 - 32)); b2 -= b1;
   tmp = b1 ^ b0; b1 = (tmp >> 58) | (tmp << (64 - 58)); b0 -= b1;
   tmp = b3 ^ b2; b3 = (tmp >> 22) | (tmp << (64 - 22)); b2 -= b3;
   tmp = b3 ^ b0; b3 = (tmp >> 46) | (tmp << (64 - 46)); b0 -= b3;
   tmp = b1 ^ b2; b1 = (tmp >> 12) | (tmp << (64 - 12)); b2 -= b1;
   tmp = b1 ^ b0; b1 = (tmp >> 25) | (tmp << (64 - 25)); b0 -= b1 + k1; b1 -= k2 + t2;
   tmp = b3 ^ b2; b3 = (tmp >> 33) | (tmp << (64 - 33)); b2 -= b3 + k3 + t0; b3 -= k4 + 11;
   tmp = b3 ^ b0; b3 = (tmp >> 5) | (tmp << (64 - 5)); b0 -= b3;
   tmp = b1 ^ b2; b1 = (tmp >> 37) | (tmp << (64 - 37)); b2 -= b1;
   tmp = b1 ^ b0; b1 = (tmp >> 23) | (tmp << (64 - 23)); b0 -= b1;
   tmp = b3 ^ b2; b3 = (tmp >> 40) | (tmp << (64 - 40)); b2 -= b3;
   tmp = b3 ^ b0; b3 = (tmp >> 52) | (tmp << (64 - 52)); b0 -= b3;
   tmp = b1 ^ b2; b1 = (tmp >> 57) | (tmp << (64 - 57)); b2 -= b1;
   tmp = b1 ^ b0; b1 = (tmp >> 14) | (tmp << (64 - 14)); b0 -= b1 + k0; b1 -= k1 + t1;
   tmp = b3 ^ b2; b3 = (tmp >> 16) | (tmp << (64 - 16)); b2 -= b3 + k2 + t2; b3 -= k3 + 10;
   tmp = b3 ^ b0; b3 = (tmp >> 32) | (tmp << (64 - 32)); b0 -= b3;
   tmp = b1 ^ b2; b1 = (tmp >> 32) | (tmp << (64 - 32)); b2 -= b1;
   tmp = b1 ^ b0; b1 = (tmp >> 58) | (tmp << (64 - 58)); b0 -= b1;
   tmp = b3 ^ b2; b3 = (tmp >> 22) | (tmp << (64 - 22)); b2 -= b3;
   tmp = b3 ^ b0; b3 = (tmp >> 46) | (tmp << (64 - 46)); b0 -= b3;
   tmp = b1 ^ b2; b1 = (tmp >> 12) | (tmp << (64 - 12)); b2 -= b1;
   tmp = b1 ^ b0; b1 = (tmp >> 25) | (tmp << (64 - 25)); b0 -= b1 + k4; b1 -= k0 + t0;
   tmp = b3 ^ b2; b3 = (tmp >> 33) | (tmp << (64 - 33)); b2 -= b3 + k1 + t1; b3 -= k2 + 9;
   tmp = b3 ^ b0; b3 = (tmp >> 5) | (tmp << (64 - 5)); b0 -= b3;
   tmp = b1 ^ b2; b1 = (tmp >> 37) | (tmp << (64 - 37)); b2 -= b1;
   tmp = b1 ^ b0; b1 = (tmp >> 23) | (tmp << (64 - 23)); b0 -= b1;
   tmp = b3 ^ b2; b3 = (tmp >> 40) | (tmp << (64 - 40)); b2 -= b3;
   tmp = b3 ^ b0; b3 = (tmp >> 52) | (tmp << (64 - 52)); b0 -= b3;
   tmp = b1 ^ b2; b1 = (tmp >> 57) | (tmp << (64 - 57)); b2 -= b1;
   tmp = b1 ^ b0; b1 = (tmp >> 14) | (tmp << (64 - 14)); b0 -= b1 + k3; b1 -= k4 + t2;
   tmp = b3 ^ b2; b3 = (tmp >> 16) | (tmp << (64 - 16)); b2 -= b3 + k0 + t0; b3 -= k1 + 8;
   tmp = b3 ^ b0; b3 = (tmp >> 32) | (tmp << (64 - 32)); b0 -= b3;
   tmp = b1 ^ b2; b1 = (tmp >> 32) | (tmp << (64 - 32)); b2 -= b1;
   tmp = b1 ^ b0; b1 = (tmp >> 58) | (tmp << (64 - 58)); b0 -= b1;
   tmp = b3 ^ b2; b3 = (tmp >> 22) | (tmp << (64 - 22)); b2 -= b3;
   tmp = b3 ^ b0; b3 = (tmp >> 46) | (tmp << (64 - 46)); b0 -= b3;
   tmp = b1 ^ b2; b1 = (tmp >> 12) | (tmp << (64 - 12)); b2 -= b1;
   tmp = b1 ^ b0; b1 = (tmp >> 25) | (tmp << (64 - 25)); b0 -= b1 + k2; b1 -= k3 + t1;
   tmp = b3 ^ b2; b3 = (tmp >> 33) | (tmp << (64 - 33)); b2 -= b3 + k4 + t2; b3 -= k0 + 7;
   tmp = b3 ^ b0; b3 = (tmp >> 5) | (tmp << (64 - 5)); b0 -= b3;
   tmp = b1 ^ b2; b1 = (tmp >> 37) | (tmp << (64 - 37)); b2 -= b1;
   tmp = b1 ^ b0; b1 = (tmp >> 23) | (tmp << (64 - 23)); b0 -= b1;
   tmp = b3 ^ b2; b3 = (tmp >> 40) | (tmp << (64 - 40)); b2 -= b3;
   tmp = b3 ^ b0; b3 = (tmp >> 52) | (tmp << (64 - 52)); b0 -= b3;
   tmp = b1 ^ b2; b1 = (tmp >> 57) | (tmp << (64 - 57)); b2 -= b1;
   tmp = b1 ^ b0; b1 = (tmp >> 14) | (tmp << (64 - 14)); b0 -= b1 + k1; b1 -= k2 + t0;
   tmp = b3 ^ b2; b3 = (tmp >> 16) | (tmp << (64 - 16)); b2 -= b3 + k3 + t1; b3 -= k4 + 6;
   tmp = b3 ^ b0; b3 = (tmp >> 32) | (tmp << (64 - 32)); b0 -= b3;
   tmp = b1 ^ b2; b1 = (tmp >> 32) | (tmp << (64 - 32)); b2 -= b1;
   tmp = b1 ^ b0; b1 = (tmp >> 58) | (tmp << (64 - 58)); b0 -= b1;
   tmp = b3 ^ b2; b3 = (tmp >> 22) | (tmp << (64 - 22)); b2 -= b3;
   tmp = b3 ^ b0; b3 = (tmp >> 46) | (tmp << (64 - 46)); b0 -= b3;
   tmp = b1 ^ b2; b1 = (tmp >> 12) | (tmp << (64 - 12)); b2 -= b1;
   tmp = b1 ^ b0; b1 = (tmp >> 25) | (tmp << (64 - 25)); b0 -= b1 + k0; b1 -= k1 + t2;
   tmp = b3 ^ b2; b3 = (tmp >> 33) | (tmp << (64 - 33)); b2 -= b3 + k2 + t0; b3 -= k3 + 5;
   tmp = b3 ^ b0; b3 = (tmp >> 5) | (tmp << (64 - 5)); b0 -= b3;
   tmp = b1 ^ b2; b1 = (tmp >> 37) | (tmp << (64 - 37)); b2 -= b1;
   tmp = b1 ^ b0; b1 = (tmp >> 23) | (tmp << (64 - 23)); b0 -= b1;
   tmp = b3 ^ b2; b3 = (tmp >> 40) | (tmp << (64 - 40)); b2 -= b3;
   tmp = b3 ^ b0; b3 = (tmp >> 52) | (tmp << (64 - 52)); b0 -= b3;
   tmp = b1 ^ b2; b1 = (tmp >> 57) | (tmp << (64 - 57)); b2 -= b1;
   tmp = b1 ^ b0; b1 = (tmp >> 14) | (tmp << (64 - 14)); b0 -= b1 + k4; b1 -= k0 + t1;
   tmp = b3 ^ b2; b3 = (tmp >> 16) | (tmp << (64 - 16)); b2 -= b3 + k1 + t2; b3 -= k2 + 4;
   tmp = b3 ^ b0; b3 = (tmp >> 32) | (tmp << (64 - 32)); b0 -= b3;
   tmp = b1 ^ b2; b1 = (tmp >> 32) | (tmp << (64 - 32)); b2 -= b1;
   tmp = b1 ^ b0; b1 = (tmp >> 58) | (tmp << (64 - 58)); b0 -= b1;
   tmp = b3 ^ b2; b3 = (tmp >> 22) | (tmp << (64 - 22)); b2 -= b3;
   tmp = b3 ^ b0; b3 = (tmp >> 46) | (tmp << (64 - 46)); b0 -= b3;
   tmp = b1 ^ b2; b1 = (tmp >> 12) | (tmp << (64 - 12)); b2 -= b1;
   tmp = b1 ^ b0; b1 = (tmp >> 25) | (tmp << (64 - 25)); b0 -= b1 + k3; b1 -= k4 + t0;
   tmp = b3 ^ b2; b3 = (tmp >> 33) | (tmp << (64 - 33)); b2 -= b3 + k0 + t1; b3 -= k1 + 3;
   tmp = b3 ^ b0; b3 = (tmp >> 5) | (tmp << (64 - 5)); b0 -= b3;
   tmp = b1 ^ b2; b1 = (tmp >> 37) | (tmp << (64 - 37)); b2 -= b1;
   tmp = b1 ^ b0; b1 = (tmp >> 23) | (tmp << (64 - 23)); b0 -= b1;
   tmp = b3 ^ b2; b3 = (tmp >> 40) | (tmp << (64 - 40)); b2 -= b3;
   tmp = b3 ^ b0; b3 = (tmp >> 52) | (tmp << (64 - 52)); b0 -= b3;
   tmp = b1 ^ b2; b1 = (tmp >> 57) | (tmp << (64 - 57)); b2 -= b1;
   tmp = b1 ^ b0; b1 = (tmp >> 14) | (tmp << (64 - 14)); b0 -= b1 + k2; b1 -= k3 + t2;
   tmp = b3 ^ b2; b3 = (tmp >> 16) | (tmp << (64 - 16)); b2 -= b3 + k4 + t0; b3 -= k0 + 2;
   tmp = b3 ^ b0; b3 = (tmp >> 32) | (tmp << (64 - 32)); b0 -= b3;
   tmp = b1 ^ b2; b1 = (tmp >> 32) | (tmp << (64 - 32)); b2 -= b1;
   tmp = b1 ^ b0; b1 = (tmp >> 58) | (tmp << (64 - 58)); b0 -= b1;
   tmp = b3 ^ b2; b3 = (tmp >> 22) | (tmp << (64 - 22)); b2 -= b3;
   tmp = b3 ^ b0; b3 = (tmp >> 46) | (tmp << (64 - 46)); b0 -= b3;
   tmp = b1 ^ b2; b1 = (tmp >> 12) | (tmp << (64 - 12)); b2 -= b1;
   tmp = b1 ^ b0; b1 = (tmp >> 25) | (tmp << (64 - 25)); b0 -= b1 + k1; b1 -= k2 + t1;
   tmp = b3 ^ b2; b3 = (tmp >> 33) | (tmp << (64 - 33)); b2 -= b3 + k3 + t2; b3 -= k4 + 1;
   tmp = b3 ^ b0; b3 = (tmp >> 5) | (tmp << (64 - 5)); b0 -= b3;
   tmp = b1 ^ b2; b1 = (tmp >> 37) | (tmp << (64 - 37)); b2 -= b1;
   tmp = b1 ^ b0; b1 = (tmp >> 23) | (tmp << (64 - 23)); b0 -= b1;
   tmp = b3 ^ b2; b3 = (tmp >> 40) | (tmp << (64 - 40)); b2 -= b3;
   tmp = b3 ^ b0; b3 = (tmp >> 52) | (tmp << (64 - 52)); b0 -= b3;
   tmp = b1 ^ b2; b1 = (tmp >> 57) | (tmp << (64 - 57)); b2 -= b1;
   tmp = b1 ^ b0; b1 = (tmp >> 14) | (tmp << (64 - 14)); b0 -= b1 + k0; b1 -= k1 + t0;
   tmp = b3 ^ b2; b3 = (tmp >> 16) | (tmp << (64 - 16)); b2 -= b3 + k2 + t1; b3 -= k3;

   output[0] = b0;
   output[1] = b1;
   output[2] = b2;
   output[3] = b3;
}

