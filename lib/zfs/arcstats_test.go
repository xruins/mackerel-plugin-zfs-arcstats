package zfs_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xruins/mackerel-plugin-zfs-arcstats/lib/zfs"
)

func TestGet(t *testing.T) {
	got := wantArcStats.Get("hits")
	want := &zfs.Row{Name: "hits", Type: 4, Value: 174896434}
	assert.Equal(t, want, got)
}

func TestParseArcStats(t *testing.T) {
	f := bytes.NewReader([]byte(arcStatsRaw))
	want := wantArcStats
	got, err := zfs.ParseArcStats(f)

	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

var wantArcStats = &zfs.ArcStats{}

func init() {
	wantArcStats.SetRows(map[string]*zfs.Row{
		"hits":                           &zfs.Row{Name: "hits", Type: 4, Value: 174896434},
		"misses":                         &zfs.Row{Name: "misses", Type: 4, Value: 18851609},
		"demand_data_hits":               &zfs.Row{Name: "demand_data_hits", Type: 4, Value: 94794984},
		"demand_data_misses":             &zfs.Row{Name: "demand_data_misses", Type: 4, Value: 2513667},
		"demand_metadata_hits":           &zfs.Row{Name: "demand_metadata_hits", Type: 4, Value: 79908225},
		"demand_metadata_misses":         &zfs.Row{Name: "demand_metadata_misses", Type: 4, Value: 632783},
		"prefetch_data_hits":             &zfs.Row{Name: "prefetch_data_hits", Type: 4, Value: 123704},
		"prefetch_data_misses":           &zfs.Row{Name: "prefetch_data_misses", Type: 4, Value: 15607210},
		"prefetch_metadata_hits":         &zfs.Row{Name: "prefetch_metadata_hits", Type: 4, Value: 69521},
		"prefetch_metadata_misses":       &zfs.Row{Name: "prefetch_metadata_misses", Type: 4, Value: 97949},
		"mru_hits":                       &zfs.Row{Name: "mru_hits", Type: 4, Value: 30145157},
		"mru_ghost_hits":                 &zfs.Row{Name: "mru_ghost_hits", Type: 4, Value: 111579},
		"mfu_hits":                       &zfs.Row{Name: "mfu_hits", Type: 4, Value: 144648513},
		"mfu_ghost_hits":                 &zfs.Row{Name: "mfu_ghost_hits", Type: 4, Value: 373974},
		"deleted":                        &zfs.Row{Name: "deleted", Type: 4, Value: 18062831},
		"mutex_miss":                     &zfs.Row{Name: "mutex_miss", Type: 4, Value: 5415},
		"access_skip":                    &zfs.Row{Name: "access_skip", Type: 4, Value: 2},
		"evict_skip":                     &zfs.Row{Name: "evict_skip", Type: 4, Value: 1145},
		"evict_not_enough":               &zfs.Row{Name: "evict_not_enough", Type: 4, Value: 700},
		"evict_l2_cached":                &zfs.Row{Name: "evict_l2_cached", Type: 4, Value: 0},
		"evict_l2_eligible":              &zfs.Row{Name: "evict_l2_eligible", Type: 4, Value: 2255312710144},
		"evict_l2_eligible_mfu":          &zfs.Row{Name: "evict_l2_eligible_mfu", Type: 4, Value: 253438624768},
		"evict_l2_eligible_mru":          &zfs.Row{Name: "evict_l2_eligible_mru", Type: 4, Value: 2001874085376},
		"evict_l2_ineligible":            &zfs.Row{Name: "evict_l2_ineligible", Type: 4, Value: 183844757504},
		"evict_l2_skip":                  &zfs.Row{Name: "evict_l2_skip", Type: 4, Value: 0},
		"hash_elements":                  &zfs.Row{Name: "hash_elements", Type: 4, Value: 583035},
		"hash_elements_max":              &zfs.Row{Name: "hash_elements_max", Type: 4, Value: 760032},
		"hash_collisions":                &zfs.Row{Name: "hash_collisions", Type: 4, Value: 1303640},
		"hash_chains":                    &zfs.Row{Name: "hash_chains", Type: 4, Value: 19228},
		"hash_chain_max":                 &zfs.Row{Name: "hash_chain_max", Type: 4, Value: 4},
		"p":                              &zfs.Row{Name: "p", Type: 4, Value: 25715246144},
		"c":                              &zfs.Row{Name: "c", Type: 4, Value: 33676691456},
		"c_min":                          &zfs.Row{Name: "c_min", Type: 4, Value: 2104793216},
		"c_max":                          &zfs.Row{Name: "c_max", Type: 4, Value: 33676691456},
		"size":                           &zfs.Row{Name: "size", Type: 4, Value: 33743421984},
		"compressed_size":                &zfs.Row{Name: "compressed_size", Type: 4, Value: 28242027008},
		"uncompressed_size":              &zfs.Row{Name: "uncompressed_size", Type: 4, Value: 29424824320},
		"overhead_size":                  &zfs.Row{Name: "overhead_size", Type: 4, Value: 2112220160},
		"hdr_size":                       &zfs.Row{Name: "hdr_size", Type: 4, Value: 192501600},
		"data_size":                      &zfs.Row{Name: "data_size", Type: 4, Value: 28890273280},
		"metadata_size":                  &zfs.Row{Name: "metadata_size", Type: 4, Value: 1463973888},
		"dbuf_size":                      &zfs.Row{Name: "dbuf_size", Type: 4, Value: 743366784},
		"dnode_size":                     &zfs.Row{Name: "dnode_size", Type: 4, Value: 1853657152},
		"bonus_size":                     &zfs.Row{Name: "bonus_size", Type: 4, Value: 595791360},
		"anon_size":                      &zfs.Row{Name: "anon_size", Type: 4, Value: 311296},
		"anon_evictable_data":            &zfs.Row{Name: "anon_evictable_data", Type: 4, Value: 0},
		"anon_evictable_metadata":        &zfs.Row{Name: "anon_evictable_metadata", Type: 4, Value: 0},
		"mru_size":                       &zfs.Row{Name: "mru_size", Type: 4, Value: 20855971328},
		"mru_evictable_data":             &zfs.Row{Name: "mru_evictable_data", Type: 4, Value: 20344817664},
		"mru_evictable_metadata":         &zfs.Row{Name: "mru_evictable_metadata", Type: 4, Value: 20029440},
		"mru_ghost_size":                 &zfs.Row{Name: "mru_ghost_size", Type: 4, Value: 12823452160},
		"mru_ghost_evictable_data":       &zfs.Row{Name: "mru_ghost_evictable_data", Type: 4, Value: 7983201280},
		"mru_ghost_evictable_metadata":   &zfs.Row{Name: "mru_ghost_evictable_metadata", Type: 4, Value: 4840250880},
		"mfu_size":                       &zfs.Row{Name: "mfu_size", Type: 4, Value: 9497964544},
		"mfu_evictable_data":             &zfs.Row{Name: "mfu_evictable_data", Type: 4, Value: 6662013440},
		"mfu_evictable_metadata":         &zfs.Row{Name: "mfu_evictable_metadata", Type: 4, Value: 356352},
		"mfu_ghost_size":                 &zfs.Row{Name: "mfu_ghost_size", Type: 4, Value: 20853040640},
		"mfu_ghost_evictable_data":       &zfs.Row{Name: "mfu_ghost_evictable_data", Type: 4, Value: 4758962176},
		"mfu_ghost_evictable_metadata":   &zfs.Row{Name: "mfu_ghost_evictable_metadata", Type: 4, Value: 16094078464},
		"l2_hits":                        &zfs.Row{Name: "l2_hits", Type: 4, Value: 0},
		"l2_misses":                      &zfs.Row{Name: "l2_misses", Type: 4, Value: 0},
		"l2_prefetch_asize":              &zfs.Row{Name: "l2_prefetch_asize", Type: 4, Value: 0},
		"l2_mru_asize":                   &zfs.Row{Name: "l2_mru_asize", Type: 4, Value: 0},
		"l2_mfu_asize":                   &zfs.Row{Name: "l2_mfu_asize", Type: 4, Value: 0},
		"l2_bufc_data_asize":             &zfs.Row{Name: "l2_bufc_data_asize", Type: 4, Value: 0},
		"l2_bufc_metadata_asize":         &zfs.Row{Name: "l2_bufc_metadata_asize", Type: 4, Value: 0},
		"l2_feeds":                       &zfs.Row{Name: "l2_feeds", Type: 4, Value: 0},
		"l2_rw_clash":                    &zfs.Row{Name: "l2_rw_clash", Type: 4, Value: 0},
		"l2_read_bytes":                  &zfs.Row{Name: "l2_read_bytes", Type: 4, Value: 0},
		"l2_write_bytes":                 &zfs.Row{Name: "l2_write_bytes", Type: 4, Value: 0},
		"l2_writes_sent":                 &zfs.Row{Name: "l2_writes_sent", Type: 4, Value: 0},
		"l2_writes_done":                 &zfs.Row{Name: "l2_writes_done", Type: 4, Value: 0},
		"l2_writes_error":                &zfs.Row{Name: "l2_writes_error", Type: 4, Value: 0},
		"l2_writes_lock_retry":           &zfs.Row{Name: "l2_writes_lock_retry", Type: 4, Value: 0},
		"l2_evict_lock_retry":            &zfs.Row{Name: "l2_evict_lock_retry", Type: 4, Value: 0},
		"l2_evict_reading":               &zfs.Row{Name: "l2_evict_reading", Type: 4, Value: 0},
		"l2_evict_l1cached":              &zfs.Row{Name: "l2_evict_l1cached", Type: 4, Value: 0},
		"l2_free_on_write":               &zfs.Row{Name: "l2_free_on_write", Type: 4, Value: 0},
		"l2_abort_lowmem":                &zfs.Row{Name: "l2_abort_lowmem", Type: 4, Value: 0},
		"l2_cksum_bad":                   &zfs.Row{Name: "l2_cksum_bad", Type: 4, Value: 0},
		"l2_io_error":                    &zfs.Row{Name: "l2_io_error", Type: 4, Value: 0},
		"l2_size":                        &zfs.Row{Name: "l2_size", Type: 4, Value: 0},
		"l2_asize":                       &zfs.Row{Name: "l2_asize", Type: 4, Value: 0},
		"l2_hdr_size":                    &zfs.Row{Name: "l2_hdr_size", Type: 4, Value: 0},
		"l2_log_blk_writes":              &zfs.Row{Name: "l2_log_blk_writes", Type: 4, Value: 0},
		"l2_log_blk_avg_asize":           &zfs.Row{Name: "l2_log_blk_avg_asize", Type: 4, Value: 0},
		"l2_log_blk_asize":               &zfs.Row{Name: "l2_log_blk_asize", Type: 4, Value: 0},
		"l2_log_blk_count":               &zfs.Row{Name: "l2_log_blk_count", Type: 4, Value: 0},
		"l2_data_to_meta_ratio":          &zfs.Row{Name: "l2_data_to_meta_ratio", Type: 4, Value: 0},
		"l2_rebuild_success":             &zfs.Row{Name: "l2_rebuild_success", Type: 4, Value: 0},
		"l2_rebuild_unsupported":         &zfs.Row{Name: "l2_rebuild_unsupported", Type: 4, Value: 0},
		"l2_rebuild_io_errors":           &zfs.Row{Name: "l2_rebuild_io_errors", Type: 4, Value: 0},
		"l2_rebuild_dh_errors":           &zfs.Row{Name: "l2_rebuild_dh_errors", Type: 4, Value: 0},
		"l2_rebuild_cksum_lb_errors":     &zfs.Row{Name: "l2_rebuild_cksum_lb_errors", Type: 4, Value: 0},
		"l2_rebuild_lowmem":              &zfs.Row{Name: "l2_rebuild_lowmem", Type: 4, Value: 0},
		"l2_rebuild_size":                &zfs.Row{Name: "l2_rebuild_size", Type: 4, Value: 0},
		"l2_rebuild_asize":               &zfs.Row{Name: "l2_rebuild_asize", Type: 4, Value: 0},
		"l2_rebuild_bufs":                &zfs.Row{Name: "l2_rebuild_bufs", Type: 4, Value: 0},
		"l2_rebuild_bufs_precached":      &zfs.Row{Name: "l2_rebuild_bufs_precached", Type: 4, Value: 0},
		"l2_rebuild_log_blks":            &zfs.Row{Name: "l2_rebuild_log_blks", Type: 4, Value: 0},
		"memory_throttle_count":          &zfs.Row{Name: "memory_throttle_count", Type: 4, Value: 0},
		"memory_direct_count":            &zfs.Row{Name: "memory_direct_count", Type: 4, Value: 0},
		"memory_indirect_count":          &zfs.Row{Name: "memory_indirect_count", Type: 4, Value: 0},
		"memory_all_bytes":               &zfs.Row{Name: "memory_all_bytes", Type: 4, Value: 67353382912},
		"memory_free_bytes":              &zfs.Row{Name: "memory_free_bytes", Type: 4, Value: 25289211904},
		"memory_available_bytes":         &zfs.Row{Name: "memory_available_bytes", Type: 3, Value: 22935279488},
		"arc_no_grow":                    &zfs.Row{Name: "arc_no_grow", Type: 4, Value: 0},
		"arc_tempreserve":                &zfs.Row{Name: "arc_tempreserve", Type: 4, Value: 0},
		"arc_loaned_bytes":               &zfs.Row{Name: "arc_loaned_bytes", Type: 4, Value: 0},
		"arc_prune":                      &zfs.Row{Name: "arc_prune", Type: 4, Value: 0},
		"arc_meta_used":                  &zfs.Row{Name: "arc_meta_used", Type: 4, Value: 4849290784},
		"arc_meta_limit":                 &zfs.Row{Name: "arc_meta_limit", Type: 4, Value: 25257518592},
		"arc_dnode_limit":                &zfs.Row{Name: "arc_dnode_limit", Type: 4, Value: 2525751859},
		"arc_meta_max":                   &zfs.Row{Name: "arc_meta_max", Type: 4, Value: 6554245152},
		"arc_meta_min":                   &zfs.Row{Name: "arc_meta_min", Type: 4, Value: 16777216},
		"async_upgrade_sync":             &zfs.Row{Name: "async_upgrade_sync", Type: 4, Value: 1368421},
		"demand_hit_predictive_prefetch": &zfs.Row{Name: "demand_hit_predictive_prefetch", Type: 4, Value: 14090312},
		"demand_hit_prescient_prefetch":  &zfs.Row{Name: "demand_hit_prescient_prefetch", Type: 4, Value: 0},
		"arc_need_free":                  &zfs.Row{Name: "arc_need_free", Type: 4, Value: 0},
		"arc_sys_free":                   &zfs.Row{Name: "arc_sys_free", Type: 4, Value: 2353932416},
		"arc_raw_size":                   &zfs.Row{Name: "arc_raw_size", Type: 4, Value: 0},
		"cached_only_in_progress":        &zfs.Row{Name: "cached_only_in_progress", Type: 4, Value: 0},
		"abd_chunk_waste_size":           &zfs.Row{Name: "abd_chunk_waste_size", Type: 4, Value: 3857920},
	})
}

const arcStatsRaw = `13 1 0x01 123 33456 7402060102 66478716886640
name                            type data
hits                            4    174896434
misses                          4    18851609
demand_data_hits                4    94794984
demand_data_misses              4    2513667
demand_metadata_hits            4    79908225
demand_metadata_misses          4    632783
prefetch_data_hits              4    123704
prefetch_data_misses            4    15607210
prefetch_metadata_hits          4    69521
prefetch_metadata_misses        4    97949
mru_hits                        4    30145157
mru_ghost_hits                  4    111579
mfu_hits                        4    144648513
mfu_ghost_hits                  4    373974
deleted                         4    18062831
mutex_miss                      4    5415
access_skip                     4    2
evict_skip                      4    1145
evict_not_enough                4    700
evict_l2_cached                 4    0
evict_l2_eligible               4    2255312710144
evict_l2_eligible_mfu           4    253438624768
evict_l2_eligible_mru           4    2001874085376
evict_l2_ineligible             4    183844757504
evict_l2_skip                   4    0
hash_elements                   4    583035
hash_elements_max               4    760032
hash_collisions                 4    1303640
hash_chains                     4    19228
hash_chain_max                  4    4
p                               4    25715246144
c                               4    33676691456
c_min                           4    2104793216
c_max                           4    33676691456
size                            4    33743421984
compressed_size                 4    28242027008
uncompressed_size               4    29424824320
overhead_size                   4    2112220160
hdr_size                        4    192501600
data_size                       4    28890273280
metadata_size                   4    1463973888
dbuf_size                       4    743366784
dnode_size                      4    1853657152
bonus_size                      4    595791360
anon_size                       4    311296
anon_evictable_data             4    0
anon_evictable_metadata         4    0
mru_size                        4    20855971328
mru_evictable_data              4    20344817664
mru_evictable_metadata          4    20029440
mru_ghost_size                  4    12823452160
mru_ghost_evictable_data        4    7983201280
mru_ghost_evictable_metadata    4    4840250880
mfu_size                        4    9497964544
mfu_evictable_data              4    6662013440
mfu_evictable_metadata          4    356352
mfu_ghost_size                  4    20853040640
mfu_ghost_evictable_data        4    4758962176
mfu_ghost_evictable_metadata    4    16094078464
l2_hits                         4    0
l2_misses                       4    0
l2_prefetch_asize               4    0
l2_mru_asize                    4    0
l2_mfu_asize                    4    0
l2_bufc_data_asize              4    0
l2_bufc_metadata_asize          4    0
l2_feeds                        4    0
l2_rw_clash                     4    0
l2_read_bytes                   4    0
l2_write_bytes                  4    0
l2_writes_sent                  4    0
l2_writes_done                  4    0
l2_writes_error                 4    0
l2_writes_lock_retry            4    0
l2_evict_lock_retry             4    0
l2_evict_reading                4    0
l2_evict_l1cached               4    0
l2_free_on_write                4    0
l2_abort_lowmem                 4    0
l2_cksum_bad                    4    0
l2_io_error                     4    0
l2_size                         4    0
l2_asize                        4    0
l2_hdr_size                     4    0
l2_log_blk_writes               4    0
l2_log_blk_avg_asize            4    0
l2_log_blk_asize                4    0
l2_log_blk_count                4    0
l2_data_to_meta_ratio           4    0
l2_rebuild_success              4    0
l2_rebuild_unsupported          4    0
l2_rebuild_io_errors            4    0
l2_rebuild_dh_errors            4    0
l2_rebuild_cksum_lb_errors      4    0
l2_rebuild_lowmem               4    0
l2_rebuild_size                 4    0
l2_rebuild_asize                4    0
l2_rebuild_bufs                 4    0
l2_rebuild_bufs_precached       4    0
l2_rebuild_log_blks             4    0
memory_throttle_count           4    0
memory_direct_count             4    0
memory_indirect_count           4    0
memory_all_bytes                4    67353382912
memory_free_bytes               4    25289211904
memory_available_bytes          3    22935279488
arc_no_grow                     4    0
arc_tempreserve                 4    0
arc_loaned_bytes                4    0
arc_prune                       4    0
arc_meta_used                   4    4849290784
arc_meta_limit                  4    25257518592
arc_dnode_limit                 4    2525751859
arc_meta_max                    4    6554245152
arc_meta_min                    4    16777216
async_upgrade_sync              4    1368421
demand_hit_predictive_prefetch  4    14090312
demand_hit_prescient_prefetch   4    0
arc_need_free                   4    0
arc_sys_free                    4    2353932416
arc_raw_size                    4    0
cached_only_in_progress         4    0
abd_chunk_waste_size            4    3857920`
