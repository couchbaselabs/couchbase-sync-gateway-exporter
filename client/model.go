package client

// Metrics JSON representation from /_expvar
type Metrics struct {
	Syncgateway struct {
		Global struct {
			ResourceUtilization struct {
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
			} `json:"resource_utilization"`
		} `json:"global"`
		PerDb map[string]struct {
			Cache struct {
				ChanCacheActiveRevs    int `json:"chan_cache_active_revs"`
				ChanCacheHits          int `json:"chan_cache_hits"`
				ChanCacheMaxEntries    int `json:"chan_cache_max_entries"`
				ChanCacheMisses        int `json:"chan_cache_misses"`
				ChanCacheNumChannels   int `json:"chan_cache_num_channels"`
				ChanCacheRemovalRevs   int `json:"chan_cache_removal_revs"`
				ChanCacheTombstoneRevs int `json:"chan_cache_tombstone_revs"`
				NumSkippedSeqs         int `json:"num_skipped_seqs"`
				RevCacheHits           int `json:"rev_cache_hits"`
				RevCacheMisses         int `json:"rev_cache_misses"`
			} `json:"cache"`
			CblReplicationPull struct {
				AttachmentPullBytes         int `json:"attachment_pull_bytes"`
				AttachmentPullCount         int `json:"attachment_pull_count"`
				MaxPending                  int `json:"max_pending"`
				NumPullReplActiveContinuous int `json:"num_pull_repl_active_continuous"`
				NumPullReplActiveOneShot    int `json:"num_pull_repl_active_one_shot"`
				NumPullReplCaughtUp         int `json:"num_pull_repl_caught_up"`
				NumPullReplSinceZero        int `json:"num_pull_repl_since_zero"`
				NumPullReplTotalContinuous  int `json:"num_pull_repl_total_continuous"`
				NumPullReplTotalOneShot     int `json:"num_pull_repl_total_one_shot"`
				RequestChangesCount         int `json:"request_changes_count"`
				RequestChangesTime          int `json:"request_changes_time"`
				RevProcessingTime           int `json:"rev_processing_time"`
				RevSendCount                int `json:"rev_send_count"`
				RevSendLatency              int `json:"rev_send_latency"`
			} `json:"cbl_replication_pull"`
			CblReplicationPush struct {
				AttachmentPushBytes int `json:"attachment_push_bytes"`
				AttachmentPushCount int `json:"attachment_push_count"`
				ConflictWriteCount  int `json:"conflict_write_count"`
				DocPushCount        int `json:"doc_push_count"`
				ProposeChangeCount  int `json:"propose_change_count"`
				ProposeChangeTime   int `json:"propose_change_time"`
				SyncFunctionCount   int `json:"sync_function_count"`
				SyncFunctionTime    int `json:"sync_function_time"`
				WriteProcessingTime int `json:"write_processing_time"`
			} `json:"cbl_replication_push"`
			Database struct {
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
			} `json:"database"`
			GsiViews struct {
				AccessCount     int `json:"access_count"`
				RoleAccessCount int `json:"roleAccess_count"`
			} `json:"gsi_views"`
			Security struct {
				AuthFailedCount  int `json:"auth_failed_count"`
				AuthSuccessCount int `json:"auth_success_count"`
				NumAccessErrors  int `json:"num_access_errors"`
				NumDocsRejected  int `json:"num_docs_rejected"`
				TotalAuthTime    int `json:"total_auth_time"`
			} `json:"security"`
			SharedBucketImport struct {
				ImportCount          int   `json:"import_count"`
				ImportErrorCount     int   `json:"import_error_count"`
				ImportProcessingTime int64 `json:"import_processing_time"`
			} `json:"shared_bucket_import"`
		} `json:"per_db"`
		PerReplication struct {
			// XXX: nothing ever appears here
		} `json:"per_replication"`
	} `json:"syncgateway"`
}
