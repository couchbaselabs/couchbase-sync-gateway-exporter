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
	SgrDocsCheckedSent               int `json:"sgr_docs_checked_sent"`
	SgrNumAttachmentBytesTransferred int `json:"sgr_num_attachment_bytes_transferred"`
	SgrNumAttachmentsTransferred     int `json:"sgr_num_attachments_transferred"`
	SgrNumDocsFailedToPush           int `json:"sgr_num_docs_failed_to_push"`
	SgrNumDocsPushed                 int `json:"sgr_num_docs_pushed"`
}

// ResourceUtilization stats
type ResourceUtilization struct {
	AdminNetBytesRecv            int     `json:"admin_net_bytes_recv"`
	AdminNetBytesSent            int     `json:"admin_net_bytes_sent"`
	ErrorCount                   int     `json:"error_count"`
	GoMemstatsHeapalloc          int     `json:"go_memstats_heapalloc"`
	GoMemstatsHeapidle           int     `json:"go_memstats_heapidle"`
	GoMemstatsHeapinuse          int     `json:"go_memstats_heapinuse"`
	GoMemstatsHeapreleased       int     `json:"go_memstats_heapreleased"`
	GoMemstatsPausetotalns       int     `json:"go_memstats_pausetotalns"`
	GoMemstatsStackinuse         int     `json:"go_memstats_stackinuse"`
	GoMemstatsStacksys           int     `json:"go_memstats_stacksys"`
	GoMemstatsSys                int     `json:"go_memstats_sys"`
	GoroutinesHighWatermark      int     `json:"goroutines_high_watermark"`
	NumGoroutines                int     `json:"num_goroutines"`
	ProcessCPUPercentUtilization float64 `json:"process_cpu_percent_utilization"`
	ProcessMemoryResident        float64 `json:"process_memory_resident"`
	PubNetBytesRecv              int     `json:"pub_net_bytes_recv"`
	PubNetBytesSent              int     `json:"pub_net_bytes_sent"`
	SystemMemoryTotal            float64 `json:"system_memory_total"`
	WarnCount                    int     `json:"warn_count"`
}

// Cache stats per db
type Cache struct {
	ChanCacheActiveRevs     int `json:"chan_cache_active_revs"`
	ChanCacheHits           int `json:"chan_cache_hits"`
	ChanCacheMaxEntries     int `json:"chan_cache_max_entries"`
	ChanCacheMisses         int `json:"chan_cache_misses"`
	ChanCacheNumChannels    int `json:"chan_cache_num_channels"`
	ChanCachePendingQueries int `json:"chan_cache_pending_queries"`
	ChanCacheRemovalRevs    int `json:"chan_cache_removal_revs"`
	ChanCacheTombstoneRevs  int `json:"chan_cache_tombstone_revs"`
	NumSkippedSeqs          int `json:"num_skipped_seqs"`
	RevCacheHits            int `json:"rev_cache_hits"`
	RevCacheMisses          int `json:"rev_cache_misses"`
}

// CblReplicationPull stats per db
type CblReplicationPull struct {
	AttachmentPullBytes         int `json:"attachment_pull_bytes"`
	AttachmentPullCount         int `json:"attachment_pull_count"`
	MaxPending                  int `json:"max_pending"`
	NumPullReplActiveContinuous int `json:"num_pull_repl_active_continuous"`
	NumPullReplActiveOneShot    int `json:"num_pull_repl_active_one_shot"`
	NumPullReplCaughtUp         int `json:"num_pull_repl_caught_up"`
	NumPullReplSinceZero        int `json:"num_pull_repl_since_zero"`
	NumPullReplTotalContinuous  int `json:"num_pull_repl_total_continuous"`
	NumPullReplTotalOneShot     int `json:"num_pull_repl_total_one_shot"`
	NumReplicationsActive       int `json:"num_replications_active"`
	RequestChangesCount         int `json:"request_changes_count"`
	RequestChangesTime          int `json:"request_changes_time"`
	RevProcessingTime           int `json:"rev_processing_time"`
	RevSendCount                int `json:"rev_send_count"`
	RevSendLatency              int `json:"rev_send_latency"`
}

// CblReplicationPush stats per db
type CblReplicationPush struct {
	AttachmentPushBytes int `json:"attachment_push_bytes"`
	AttachmentPushCount int `json:"attachment_push_count"`
	ConflictWriteCount  int `json:"conflict_write_count"`
	DocPushCount        int `json:"doc_push_count"`
	ProposeChangeCount  int `json:"propose_change_count"`
	ProposeChangeTime   int `json:"propose_change_time"`
	SyncFunctionCount   int `json:"sync_function_count"`
	SyncFunctionTime    int `json:"sync_function_time"`
	WriteProcessingTime int `json:"write_processing_time"`
}

// Database stats per db
type Database struct {
	AbandonedSeqs           int `json:"abandoned_seqs"`
	Crc32CMatchCount        int `json:"crc32c_match_count"`
	DcpCachingCount         int `json:"dcp_caching_count"`
	DcpCachingTime          int `json:"dcp_caching_time"`
	DcpReceivedCount        int `json:"dcp_received_count"`
	DcpReceivedTime         int `json:"dcp_received_time"`
	DocReadsBytesBlip       int `json:"doc_reads_bytes_blip"`
	DocWritesBytes          int `json:"doc_writes_bytes"`
	DocWritesBytesBlip      int `json:"doc_writes_bytes_blip"`
	NumDocReadsBlip         int `json:"num_doc_reads_blip"`
	NumDocReadsRest         int `json:"num_doc_reads_rest"`
	NumDocWrites            int `json:"num_doc_writes"`
	NumReplicationsActive   int `json:"num_replications_active"`
	NumReplicationsTotal    int `json:"num_replications_total"`
	SequenceGetCount        int `json:"sequence_get_count"`
	SequenceReleasedCount   int `json:"sequence_released_count"`
	SequenceReservedCount   int `json:"sequence_reserved_count"`
	WarnChannelsPerDocCount int `json:"warn_channels_per_doc_count"`
	WarnGrantsPerDocCount   int `json:"warn_grants_per_doc_count"`
	WarnXattrSizeCount      int `json:"warn_xattr_size_count"`
}

// GsiViews stats per db
type GsiViews struct {
	AccessCount     int `json:"access_count"`
	RoleAccessCount int `json:"roleAccess_count"`
	ChannelsCount   int `json:"channels_count"`
	AllDocsCount    int `json:"allDocs_count"`
	PrincipalsCount int `json:"principals_count"`
	ResyncCount     int `json:"resync_count"`
	SequencesCount  int `json:"sequences_count"`
	SessionsCount   int `json:"sessions_count"`
	TombstonesCount int `json:"tombstones_count"`
}

// Security stats per db
type Security struct {
	AuthFailedCount  int `json:"auth_failed_count"`
	AuthSuccessCount int `json:"auth_success_count"`
	NumAccessErrors  int `json:"num_access_errors"`
	NumDocsRejected  int `json:"num_docs_rejected"`
	TotalAuthTime    int `json:"total_auth_time"`
}

// SharedBucketImport stats per db
type SharedBucketImport struct {
	ImportCount          int   `json:"import_count"`
	ImportErrorCount     int   `json:"import_error_count"`
	ImportProcessingTime int64 `json:"import_processing_time"`
}

type DeltaSync struct {
	DeltaCacheHit             int `json:"delta_cache_hit"`
	DeltaCacheMiss            int `json:"delta_cache_miss"`
	DeltaPullReplicationCount int `json:"delta_pull_replication_count"`
	DeltaPushDocCount         int `json:"delta_push_doc_count"`
	DeltasRequested           int `json:"deltas_requested"`
	DeltasSent                int `json:"deltas_sent"`
}
