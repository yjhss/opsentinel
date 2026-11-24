#!/usr/bin/env bash
set -uo pipefail

failed=0

check_tool() {
  local tool="$1"

  if command -v "$tool" >/dev/null 2>&1; then
    echo "[PASS] $tool"
  else
    echo "[FAIL] $tool 설치 필요"
    failed=1
  fi
}

echo "=== Opsentinel 개발 환경 검증 ==="

if [[ "$(uname -s)" != "Darwin" ]]; then
  echo "[FAIL] macOS 환경이 아닙니다."
  exit 1
fi

echo "CPU: $(uname -m)"

for tool in brew go git gh jq docker; do
  check_tool "$tool"
done

if command -v docker >/dev/null 2>&1; then
  if docker info >/dev/null 2>&1; then
    echo "[PASS] Docker Engine"
  else
    echo "[FAIL] Docker Engine에 연결할 수 없습니다."
    failed=1
  fi

  if docker compose version >/dev/null 2>&1; then
    echo "[PASS] Docker Compose"
  else
    echo "[FAIL] Docker Compose 사용 불가"
    failed=1
  fi
fi

if [[ "$failed" -eq 0 ]]; then
  echo "환경 검증 성공"
else
  echo "환경 검증 실패"
fi

exit "$failed"
