package job

import "testing"

func TestSampleJobs(t *testing.T) {
	jobs := SampleJobs("v0.1.0")

	if len(jobs) != 3 {
		t.Fatalf("예상 작업 3건, 실제 %d건", len(jobs))
	}

	seen := make(map[string]bool)

	for _, j := range jobs {
		if j.ID == "" {
			t.Fatal("작업 ID가 비어 있습니다")
		}

		if seen[j.ID] {
			t.Fatalf("중복 작업 ID: %s", j.ID)
		}
		seen[j.ID] = true

		if j.ReleaseVersion != "v0.1.0" {
			t.Fatalf("작업 버전 불일치: %s", j.ReleaseVersion)
		}
	}
}
