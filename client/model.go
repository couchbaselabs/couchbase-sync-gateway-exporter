package client

// Metrics JSON representation from /_expvar
type Metrics struct {
	Syncgateway struct {
		Global struct {
			ResourceUtilization ResourceUtilization `json:"resource_utilization"`
		} `json:"global"`
		PerDb map[string]struct {
			Cache              Cache              `json:"cache"`
			CblReplicationPull CblReplicationPull `json:"cbl_replication_pull"`
			CblReplicationPush CblReplicationPush `json:"cbl_replication_push"`
			Database           Database           `json:"database"`
			GsiViews           GsiViews           `json:"gsi_views"`
			Security           Security           `json:"security"`
			SharedBucketImport SharedBucketImport `json:"shared_bucket_import"`
			DeltaSync          DeltaSync          `json:"delta_sync"`
		} `json:"per_db"`
		PerReplication map[string]Replication `json:"per_replication"`
	} `json:"syncgateway"`
}

// Replication stats for each replication config
type Replication struct {
	SgrDocsCheckedSent               float64 `json:"sgr_docs_checked_sent"`
	SgrNumAttachmentBytesTransferred float64 `json:"sgr_num_attachment_bytes_transferred"`
	SgrNumAttachmentsTransferred     float64 `json:"sgr_num_attachments_transferred"`
	SgrNumDocsFailedToPush           float64 `json:"sgr_num_docs_failed_to_push"`
	SgrNumDocsPushed                 float64 `json:"sgr_num_docs_pushed"`
}

// ResourceUtilization stats
type ResourceUtilization struct {
	AdminNetBytesRecv            float64 `json:"admin_net_bytes_recv"`
	AdminNetBytesSent            float64 `json:"admin_net_bytes_sent"`
	ErrorCount                   float64 `json:"error_count"`
	GoMemstatsHeapalloc          float64 `json:"go_memstats_heapalloc"`
	GoMemstatsHeapidle           float64 `json:"go_memstats_heapidle"`
	GoMemstatsHeapinuse          float64 `json:"go_memstats_heapinuse"`
	GoMemstatsHeapreleased       float64 `json:"go_memstats_heapreleased"`
	GoMemstatsPausetotalns       float64 `json:"go_memstats_pausetotalns"`
	GoMemstatsStackinuse         float64 `json:"go_memstats_stackinuse"`
	GoMemstatsStacksys           float64 `json:"go_memstats_stacksys"`
	GoMemstatsSys                float64 `json:"go_memstats_sys"`
	GoroutinesHighWatermark      float64 `json:"goroutines_high_watermark"`
	NumGoroutines                float64 `json:"num_goroutines"`
	ProcessCPUPercentUtilization float64 `json:"process_cpu_percent_utilization"`
	ProcessMemoryResident        float64 `json:"process_memory_resident"`
	PubNetBytesRecv              float64 `json:"pub_net_bytes_recv"`
	PubNetBytesSent              float64 `json:"pub_net_bytes_sent"`
	SystemMemoryTotal            float64 `json:"system_memory_total"`
	WarnCount                    float64 `json:"warn_count"`
}

// Cache stats per db
type Cache struct {
	ChanCacheActiveRevs     float64 `json:"chan_cache_active_revs"`
	ChanCacheHits           float64 `json:"chan_cache_hits"`
	ChanCacheMaxEntries     float64 `json:"chan_cache_max_entries"`
	ChanCacheMisses         float64 `json:"chan_cache_misses"`
	ChanCacheNumChannels    float64 `json:"chan_cache_num_channels"`
	ChanCachePendingQueries float64 `json:"chan_cache_pending_queries"`
	ChanCacheRemovalRevs    float64 `json:"chan_cache_removal_revs"`
	ChanCacheTombstoneRevs  float64 `json:"chan_cache_tombstone_revs"`
	NumSkippedSeqs          float64 `json:"num_skipped_seqs"`
	RevCacheHits            float64 `json:"rev_cache_hits"`
	RevCacheMisses          float64 `json:"rev_cache_misses"`
}

// CblReplicationPull stats per db
type CblReplicationPull struct {
	AttachmentPullBytes         float64 `json:"attachment_pull_bytes"`
	AttachmentPullCount         float64 `json:"attachment_pull_count"`
	MaxPending                  float64 `json:"max_pending"`
	NumPullReplActiveContinuous float64 `json:"num_pull_repl_active_continuous"`
	NumPullReplActiveOneShot    float64 `json:"num_pull_repl_active_one_shot"`
	NumPullReplCaughtUp         float64 `json:"num_pull_repl_caught_up"`
	NumPullReplSinceZero        float64 `json:"num_pull_repl_since_zero"`
	NumPullReplTotalContinuous  float64 `json:"num_pull_repl_total_continuous"`
	NumPullReplTotalOneShot     float64 `json:"num_pull_repl_total_one_shot"`
	RequestChangesCount         float64 `json:"request_changes_count"`
	RequestChangesTime          float64 `json:"request_changes_time"`
	RevProcessingTime           float64 `json:"rev_processing_time"`
	RevSendCount                float64 `json:"rev_send_count"`
	RevSendLatency              float64 `json:"rev_send_latency"`
}

// CblReplicationPush stats per db
type CblReplicationPush struct {
	AttachmentPushBytes float64 `json:"attachment_push_bytes"`
	AttachmentPushCount float64 `json:"attachment_push_count"`
	ConflictWriteCount  float64 `json:"conflict_write_count"`
	DocPushCount        float64 `json:"doc_push_count"`
	ProposeChangeCount  float64 `json:"propose_change_count"`
	ProposeChangeTime   float64 `json:"propose_change_time"`
	SyncFunctionCount   float64 `json:"sync_function_count"`
	SyncFunctionTime    float64 `json:"sync_function_time"`
	WriteProcessingTime float64 `json:"write_processing_time"`
}

// Database stats per db
type Database struct {
	AbandonedSeqs           float64 `json:"abandoned_seqs"`
	Crc32CMatchCount        float64 `json:"crc32c_match_count"`
	DcpCachingCount         float64 `json:"dcp_caching_count"`
	DcpCachingTime          float64 `json:"dcp_caching_time"`
	DcpReceivedCount        float64 `json:"dcp_received_count"`
	DcpReceivedTime         float64 `json:"dcp_received_time"`
	DocReadsBytesBlip       float64 `json:"doc_reads_bytes_blip"`
	DocWritesBytes          float64 `json:"doc_writes_bytes"`
	DocWritesBytesBlip      float64 `json:"doc_writes_bytes_blip"`
	NumDocReadsBlip         float64 `json:"num_doc_reads_blip"`
	NumDocReadsRest         float64 `json:"num_doc_reads_rest"`
	NumDocWrites            float64 `json:"num_doc_writes"`
	NumReplicationsActive   float64 `json:"num_replications_active"`
	NumReplicationsTotal    float64 `json:"num_replications_total"`
	SequenceGetCount        float64 `json:"sequence_get_count"`
	SequenceReleasedCount   float64 `json:"sequence_released_count"`
	SequenceReservedCount   float64 `json:"sequence_reserved_count"`
	WarnChannelsPerDocCount float64 `json:"warn_channels_per_doc_count"`
	WarnGrantsPerDocCount   float64 `json:"warn_grants_per_doc_count"`
	WarnXattrSizeCount      float64 `json:"warn_xattr_size_count"`
}

// GsiViews stats per db
type GsiViews struct {
	AccessCount     float64 `json:"access_count"`
	RoleAccessCount float64 `json:"roleAccess_count"`
	ChannelsCount   float64 `json:"channels_count"`
	AllDocsCount    float64 `json:"allDocs_count"`
	PrincipalsCount float64 `json:"principals_count"`
	ResyncCount     float64 `json:"resync_count"`
	SequencesCount  float64 `json:"sequences_count"`
	SessionsCount   float64 `json:"sessions_count"`
	TombstonesCount float64 `json:"tombstones_count"`
}

// Security stats per db
type Security struct {
	AuthFailedCount  float64 `json:"auth_failed_count"`
	AuthSuccessCount float64 `json:"auth_success_count"`
	NumAccessErrors  float64 `json:"num_access_errors"`
	NumDocsRejected  float64 `json:"num_docs_rejected"`
	TotalAuthTime    float64 `json:"total_auth_time"`
}

// SharedBucketImport stats per db
type SharedBucketImport struct {
	ImportCount          float64 `json:"import_count"`
	ImportErrorCount     float64 `json:"import_error_count"`
	ImportProcessingTime float64 `json:"import_processing_time"`
}

// DeltaSync stats per db
type DeltaSync struct {
	DeltaCacheHit             float64 `json:"delta_cache_hit"`
	DeltaCacheMiss            float64 `json:"delta_cache_miss"`
	DeltaPullReplicationCount float64 `json:"delta_pull_replication_count"`
	DeltaPushDocCount         float64 `json:"delta_push_doc_count"`
	DeltasRequested           float64 `json:"deltas_requested"`
	DeltasSent                float64 `json:"deltas_sent"`
}
