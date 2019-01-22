package test

import (
	"testing"
	"github.com/lwl1989/go-spider/spider"
	"fmt"
)

func Test_RemoveScript(t *testing.T) {
	//str := "<script>d</script>dasdsad啊啊啊啊是飞洒发asadas手打<script>冬天的风儿吹着雪花</script>dsadasda"
	//str = spider.RemoveScript(str)
	//fmt.Println(str)
	//
	//str = "dasdsad啊啊啊啊是飞洒发asadas手打<script>冬天的风儿吹着雪花</script>dsadasda"
	//str = spider.RemoveScript(str)
	//fmt.Println(str)

	str := `<div id="story_body_content">
			<h2 id="story_art_title">不滿同意去蔣化 文化部長遭藝人呼巴掌</h2>
			<div id="shareBar" class="shareBar">
					<div class="shareBar__main">
		<ul class="shareBar__main--bar">
			<li class="fb"><a href="#" title="分享">分享</a></li>
			<li class="line"><a href="#" title="分享">分享</a></li>
			<li class="whatsapp">
				<a href="whatsapp://send?text=不滿同意去蔣化 文化部長遭藝人呼巴掌 | 政經大事 | 產業 | 經濟日報:https://money.udn.com/money/story/7307/3607830" onclick="window.open(&#39;whatsapp://send?text=不滿同意去蔣化 文化部長遭藝人呼巴掌 | 政經大事 | 產業 | 經濟日報:https://money.udn.com/money/story/7307/3607830&#39;, &#39;_blank&#39;, &#39;toolbar=no,scrollbars=yes,resizable=no,fullscreen=no,top=50,left=50,width=600,height=600&#39;)" data-action="share/whatsapp/share" class="whatsapp-link" target="_blank"></a>
			</li>
			<li class="discuss"><a href="#comments" title="留言">留言</a></li>
			<li class="print"><a href="#" title="列印">列印</a></li>
			<li class="save">

			</li>
			<li id="shareBar--close"><a href="javascript:void(0);"></a></li>
		</ul>
		<div class="set_font_size only_web">
	<a href="#" data-attr="-">A-</a>
	<a href="#" data-attr="+">A+</a>
</div>	</div>
	<!-- /.shareBar__main -->

				<div class="shareBar__info">
					<div class="shareBar__info--author"><span>2019-01-22 14:56</span>聯合報 記者林怡秀╱即時報導</div>
					<ul class="shareBar__info--push">
						<li class="fbsend">
							<div id="fb-root"></div>
							<div class="fb-send facebook-holder" data-href="https://money.udn.com/money/story/7307/3607830"></div>
						</li>
						<li class="fblike">
							<div id="fb-root"></div>
							<div class="fb-like facebook-holder" data-href="https://money.udn.com/money/story/7307/3607830" data-width="100" data-layout="button_count" data-action="like" data-show-faces="false" data-share="true"></div>
						</li>
						<li class="linelike">
							<a href="####" title="可將此頁面的資訊分享至LINE。"><span>用LINE傳送</span></a>
						</li>
						<li id="shareBar--open"><a href="javascript:void(0);"></a></li>
					</ul>
					<script>
						function colorboxexplan(url, w, h) { window.$.colorbox({ href: url, iframe: true, scrolling: false, width: w, height: h, opacity: 0, transition: "none" }); }
					</script>
				</div>
				<!-- /.shareBar__info -->
			</div>
			<!--/.shareBar-->
			<p></p><figure class="photo_center"><a href="https://pgw.udn.com.tw/gw/photo.php?u=https://uc.udn.com.tw/photo/2019/01/22/realtime/5829720.jpg&amp;x=0&amp;y=0&amp;sw=0&amp;sh=0&amp;sl=W&amp;fw=1050&amp;exp=3600&amp;exp=3600" rel="prettyPhoto[pp_gal]"><img src="https://pgw.udn.com.tw/gw/photo.php?u=https://uc.udn.com.tw/photo/2019/01/22/realtime/5829720.jpg&amp;x=0&amp;y=0&amp;sw=0&amp;sh=0&amp;sl=W&amp;fw=1050&amp;exp=3600&amp;exp=3600" alt="周遊和李朝永夫婦在關懷演藝人員春節餐會上致詞，台下坐著常楓、文化部長鄭麗君和文夏..." title="周遊和李朝永夫婦在關懷演藝人員春節餐會上致詞，台下坐著常楓、文化部長鄭麗君和文夏..."/></a><figcaption>周遊和李朝永夫婦在關懷演藝人員春節餐會上致詞，台下坐著常楓、文化部長鄭麗君和文夏。 記者杜建重／攝影</figcaption><div class="photo_pop"><ul><li class="facebook"><a title="facebook" href="javascript:addFacebook(&#39;img&#39;, &#39;https://pgw.udn.com.tw/gw/photo.php?u=https://uc.udn.com.tw/photo/2019/01/22/realtime/5829720.jpg&amp;x=0&amp;y=0&amp;sw=0&amp;sh=0&amp;sl=W&amp;fw=1050&amp;exp=3600&amp;exp=3600&#39;);">facebook</a></li></ul></div></figure><!-- /.photo --><p></p><p>
文化部長鄭麗君今出席資深藝人春節餐會，逐桌敬酒時竟遭資深藝人鄭惠中呼巴掌，似乎是因為不滿鄭麗君同意去蔣化，所以才動手打人，鄭麗君當場錯愕，而華視表示，鄭惠中事後有道歉，一切純屬她個人行為。</p><!--999-->												<div id="_popIn_recommend"></div>
			<script type="text/javascript">
				(function() {
					var pa = document.createElement('script'); pa.type = 'text/javascript'; pa.charset = "utf-8"; pa.async = true;
					pa.src = window.location.protocol + "//api.popin.cc/searchbox/udn_money.js";
					var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(pa, s);
				})();
			</script>
		</div>
		<!-- /#story_body_content -->
					<div id="story_tags">
									<a href="/search/tagging/1001/鄭麗君" rel="113185">鄭麗君</a>
									<a href="/search/tagging/1001/春節" rel="124306">春節</a>
							</div>
								<div>
	<!--123456789--></div>`
	str = spider.RemoveScript(str)
	str = spider.RemoveSpace(str)
	fmt.Println(str)
}
