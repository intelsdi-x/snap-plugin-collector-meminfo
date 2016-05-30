/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2015-2016 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mem

type MemMetrics struct {
	MemTotal          uint64 `json:"mem_total" mapstructure:"mem_total"`
	MemUsed           uint64 `json:"mem_used" mapstructure:"mem_used"`
	MemFree           uint64 `json:"mem_free" mapstructure:"mem_free"`
	MemAvailable      uint64 `json:"mem_available" mapstructure:"mem_available"`
	Buffers           uint64 `json:"buffers" mapstructure:"buffers"`
	Cached            uint64 `json:"cached" mapstructure:"cached"`
	SwapCached        uint64 `json:"swap_cached" mapstructure:"swap_cached"`
	Active            uint64 `json:"active" mapstructure:"active"`
	Inactive          uint64 `json:"inactive" mapstructure:"inactive"`
	ActiveAnon        uint64 `json:"active_anon" mapstructure:"active_anon"`
	InactiveAnon      uint64 `json:"inactive_anon" mapstructure:"inactive_anon"`
	ActiveFile        uint64 `json:"active_file" mapstructure:"active_file"`
	InactiveFile      uint64 `json:"inactive_file" mapstructure:"inactive_file"`
	Unevictable       uint64 `json:"unevictable" mapstructure:"unevictable"`
	Mlocked           uint64 `json:"mlocked" mapstructure:"mlocked"`
	HighTotal         uint64 `json:"high_total" mapstructure:"high_total"`
	HighFree          uint64 `json:"high_free" mapstructure:"high_free"`
	LowTotal          uint64 `json:"low_total" mapstructure:"low_total"`
	LowFree           uint64 `json:"low_free" mapstructure:"low_free"`
	MmapCopy          uint64 `json:"mmap_copy" mapstructure:"mmap_copy"`
	SwapTotal         uint64 `json:"swap_total" mapstructure:"swap_total"`
	SwapFree          uint64 `json:"swap_free" mapstructure:"swap_free"`
	Dirty             uint64 `json:"dirty" mapstructure:"dirty"`
	Writeback         uint64 `json:"writeback" mapstructure:"writeback"`
	AnonPages         uint64 `json:"anon_pages" mapstructure:"anon_pages"`
	Mapped            uint64 `json:"mapped" mapstructure:"mapped"`
	Shmem             uint64 `json:"shmem" mapstructure:"shmem"`
	Slab              uint64 `json:"slab" mapstructure:"slab"`
	SReclaimable      uint64 `json:"sreclaimable" mapstructure:"sreclaimable"`
	SUnreclaim        uint64 `json:"sunreclaim" mapstructure:"sunreclaim"`
	KernelStack       uint64 `json:"kernel_stack" mapstructure:"kernel_stack"`
	PageTables        uint64 `json:"page_tables" mapstructure:"page_tables"`
	Quicklists        uint64 `json:"quicklists" mapstructure:"quicklists"`
	NFSUnstable       uint64 `json:"nfs_unstable" mapstructure:"nfs_unstable"`
	Bounce            uint64 `json:"bounce" mapstructure:"bounce"`
	WritebackTmp      uint64 `json:"writeback_tmp" mapstructure:"writeback_tmp"`
	CommitLimit       uint64 `json:"commit_limit" mapstructure:"commit_limit"`
	CommittedAS       uint64 `json:"committed_as" mapstructure:"committed_as"`
	VmallocTotal      uint64 `json:"vmalloc_total" mapstructure:"vmalloc_total"`
	VmallocUsed       uint64 `json:"vmalloc_used" mapstructure:"vmalloc_used"`
	VmallocChunk      uint64 `json:"vmalloc_chunk" mapstructure:"vmalloc_chunk"`
	HardwareCorrupted uint64 `json:"hardware_corrupted" mapstructure:"hardware_corrupted"`
	AnonHugePages     uint64 `json:"anon_huge_pages" mapstructure:"anon_huge_pages"`
	CmaTotal          uint64 `json:"cma_total" mapstructure:"cma_total"`
	CmaFree           uint64 `json:"cma_free" mapstructure:"cma_free"`
	HugePagesTotal    uint64 `json:"huge_pages_total" mapstructure:"huge_pages_total"`
	HugePagesFree     uint64 `json:"huge_pages_free" mapstructure:"huge_pages_free"`
	HugePagesRsvd     uint64 `json:"huge_pages_rsvd" mapstructure:"huge_pages_rsvd"`
	HugePagesSurp     uint64 `json:"huge_pages_surp" mapstructure:"huge_pages_surp"`
	Hugepagesize      uint64 `json:"hugepagesize" mapstructure:"hugepagesize"`
	DirectMap4k       uint64 `json:"direct_map4k" mapstructure:"direct_map4k"`
	DirectMap4M       uint64 `json:"direct_map4m" mapstructure:"direct_map4m"`
	DirectMap2M       uint64 `json:"direct_map2m" mapstructure:"direct_map2m"`
	DirectMap1G       uint64 `json:"direct_map1g" mapstructure:"direct_map1g"`

	MemTotalPerc          float64 `json:"mem_total_perc" mapstructure:"mem_total_perc"`
	MemUsedPerc           float64 `json:"mem_used_perc" mapstructure:"mem_used_perc"`
	MemFreePerc           float64 `json:"mem_free_perc" mapstructure:"mem_free_perc"`
	MemAvailablePerc      float64 `json:"mem_available_perc" mapstructure:"mem_available_perc"`
	BuffersPerc           float64 `json:"buffers_perc" mapstructure:"buffers_perc"`
	CachedPerc            float64 `json:"cached_perc" mapstructure:"cached_perc"`
	SwapCachedPerc        float64 `json:"swap_cached_perc" mapstructure:"swap_cached_perc"`
	ActivePerc            float64 `json:"active_perc" mapstructure:"active_perc"`
	InactivePerc          float64 `json:"inactive_perc" mapstructure:"inactive_perc"`
	ActiveAnonPerc        float64 `json:"active_anon_perc" mapstructure:"active_anon_perc"`
	InactiveAnonPerc      float64 `json:"inactive_anon_perc" mapstructure:"inactive_anon_perc"`
	ActiveFilePerc        float64 `json:"active_file_perc" mapstructure:"active_file_perc"`
	InactiveFilePerc      float64 `json:"inactive_file_perc" mapstructure:"inactive_file_perc"`
	UnevictablePerc       float64 `json:"unevictable_perc" mapstructure:"unevictable_perc"`
	MlockedPerc           float64 `json:"mlocked_perc" mapstructure:"mlocked_perc"`
	HighTotalPerc         float64 `json:"high_total_perc" mapstructure:"high_total_perc"`
	HighFreePerc          float64 `json:"high_free_perc" mapstructure:"high_free_perc"`
	LowTotalPerc          float64 `json:"low_total_perc" mapstructure:"low_total_perc"`
	LowFreePerc           float64 `json:"low_free_perc" mapstructure:"low_free_perc"`
	MmapCopyPerc          float64 `json:"mmap_copy_perc" mapstructure:"mmap_copy_perc"`
	SwapTotalPerc         float64 `json:"swap_total_perc" mapstructure:"swap_total_perc"`
	SwapFreePerc          float64 `json:"swap_free_perc" mapstructure:"swap_free_perc"`
	DirtyPerc             float64 `json:"dirty_perc" mapstructure:"dirty_perc"`
	WritebackPerc         float64 `json:"writeback_perc" mapstructure:"writeback_perc"`
	AnonPagesPerc         float64 `json:"anon_pages_perc" mapstructure:"anon_pages_perc"`
	MappedPerc            float64 `json:"mapped_perc" mapstructure:"mapped_perc"`
	ShmemPerc             float64 `json:"shmem_perc" mapstructure:"shmem_perc"`
	SlabPerc              float64 `json:"slab_perc" mapstructure:"slab_perc"`
	SReclaimablePerc      float64 `json:"sreclaimable_perc" mapstructure:"sreclaimable_perc"`
	SUnreclaimPerc        float64 `json:"sunreclaim_perc" mapstructure:"sunreclaim_perc"`
	KernelStackPerc       float64 `json:"kernel_stack_perc" mapstructure:"kernel_stack_perc"`
	PageTablesPerc        float64 `json:"page_tables_perc" mapstructure:"page_tables_perc"`
	QuicklistsPerc        float64 `json:"quicklists_perc" mapstructure:"quicklists_perc"`
	NFSUnstablePerc       float64 `json:"nfs_unstable_perc" mapstructure:"nfs_unstable_perc"`
	BouncePerc            float64 `json:"bounce_perc" mapstructure:"bounce_perc"`
	WritebackTmpPerc      float64 `json:"writeback_tmp_perc" mapstructure:"writeback_tmp_perc"`
	CommitLimitPerc       float64 `json:"commit_limit_perc" mapstructure:"commit_limit_perc"`
	CommittedASPerc       float64 `json:"committed_as_perc" mapstructure:"committed_as_perc"`
	VmallocTotalPerc      float64 `json:"vmalloc_total_perc" mapstructure:"vmalloc_total_perc"`
	VmallocUsedPerc       float64 `json:"vmalloc_used_perc" mapstructure:"vmalloc_used_perc"`
	VmallocChunkPerc      float64 `json:"vmalloc_chunk_perc" mapstructure:"vmalloc_chunk_perc"`
	HardwareCorruptedPerc float64 `json:"hardware_corrupted_perc" mapstructure:"hardware_corrupted_perc"`
	AnonHugePagesPerc     float64 `json:"anon_huge_pages_perc" mapstructure:"anon_huge_pages_perc"`
	CmaTotalPerc          float64 `json:"cma_total_perc" mapstructure:"cma_total_perc"`
	CmaFreePerc           float64 `json:"cma_free_perc" mapstructure:"cma_free_perc"`
	HugePagesTotalPerc    float64 `json:"huge_pages_total_perc" mapstructure:"huge_pages_total_perc"`
	HugePagesFreePerc     float64 `json:"huge_pages_free_perc" mapstructure:"huge_pages_free_perc"`
	HugePagesRsvdPerc     float64 `json:"huge_pages_rsvd_perc" mapstructure:"huge_pages_rsvd_perc"`
	HugePagesSurpPerc     float64 `json:"huge_pages_surp_perc" mapstructure:"huge_pages_surp_perc"`
	HugepagesizePerc      float64 `json:"hugepagesize_perc" mapstructure:"hugepagesize_perc"`
	DirectMap4kPerc       float64 `json:"direct_map4k_perc" mapstructure:"direct_map4k_perc"`
	DirectMap4MPerc       float64 `json:"direct_map4m_perc" mapstructure:"direct_map4m_perc"`
	DirectMap2MPerc       float64 `json:"direct_map2m_perc" mapstructure:"direct_map2m_perc"`
	DirectMap1GPerc       float64 `json:"direct_map1g_perc" mapstructure:"direct_map1g_perc"`
}
