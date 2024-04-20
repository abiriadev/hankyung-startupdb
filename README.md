# 한국경제 스타트업 목록 스크래퍼

## 실행

```sh
go run .
```

## 구조

| 속성                 | 타입   | 설명                 |
| -------------------- | ------ | -------------------- |
| `name`               | string | 기업명               |
| `logo`               | string | 기업 로고            |
| `representative`     | string | 대표                 |
| `location`           | string | 본사 소재지          |
| `establishedAt`      | string | 설립일자             |
| `link`               | string | 홈페이지             |
| `mail`               | string | 이메일 주소          |
| `telephone`          | string | 전화번호             |
| `domain`             | string | 업종                 |
| `mainProduct`        | string | 주요사업             |
| `cLevel`             | string | C레벨 구성           |
| `employees`          | string | 인력규모             |
| `investment`         | string | 누적투자금 (억 단위) |
| `series`             | string | 투자단계             |
| `investmentOverview` | string | 투자소개서           |
| `investor`           | string | 투자사               |

## 예시

```json
{
	"name": "야놀자",
	"logo": "https://www.hankyung.com/data/service/geeks/logo/c69698bf84e31f1b8624ff4b5455e727.jpg",
	"representative": "이수진, 김종윤, 배보찬",
	"location": "서울 강남구 테헤란로 108길 42",
	"establishedAt": "2007년 2월",
	"link": "yanolja.com",
	"mail": "help@yanolja.com",
	"telephone": "16441346",
	"domain": "e커머스",
	"mainProduct": "숙박 예약 및 레저 액티비티 티켓 구매 서비스를 제공하는 글로벌 여가 플랫폼",
	"cLevel": "이수진(총괄대표), 김종윤(대표), 배보찬(CFO), 신정인(COO)",
	"employees": "1,183",
	"investment": "22,340",
	"series": "Series E",
	"investmentOverview": "",
	"investor": "소프트뱅크인베스트먼트어드바이저"
}
```
