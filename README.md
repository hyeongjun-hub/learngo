# learngo

## 사람인 사이트의 기업공고를 자동으로 추출할 수 있는 코드입니다.

설명: 
baseURL에 사람인 채용공고의 검색어를 변경하여 사용할 수 있습니다.
default는 "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=devops" 로 "devops" 검색어가 들어있습니다.
* `searchword=devops` 부분을 변경가능 


### go routine 사용
goroutine을 사용하여 최대 10개의 페이지에서 병렬적으로 scrapping 가능합니다.

### csv 파일 생성
learngo 폴더 안에 jobs.svc 파일이 생성되며 csv 리더로 확인 가능합니다.

### 테이블 구조
header : "Enterprise", "Title", "Location", "Link"
- Enterprise : 기업명
- Title : 공고명
- Location : 기업 위치
- Link : 상세 페이지 링크

### 예시사진
![image](https://user-images.githubusercontent.com/77392219/229047893-cb93f695-3245-47de-9372-8bf3c31de1e6.png)

