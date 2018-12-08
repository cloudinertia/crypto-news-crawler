package Coindesk_korea

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func Coindesk_korea_scrape_page(page int) {
	// Request the HTML page.
	res, err := http.PostForm("https://www.coindeskkorea.com/wp-content/plugins/divi_template_builder/template_builder_ajax.php", url.Values{
		"_wpnonce":         {"bf51bfeb55"},
		"_wp_http_referer": {"/news/"},
		"action":           {"more_tb_loop_archive"},
		"form_args":        {"a:7:{s:14:\"posts_per_page\";i:10;s:7:\"orderby\";s:4:\"date\";s:5:\"order\";s:4:\"desc\";s:9:\"post_type\";s:4:\"post\";s:5:\"paged\";i:0;s:6:\"offset\";i:3;s:9:\"tax_query\";a:1:{i:0;a:4:{s:8:\"taxonomy\";s:13:\"post_curation\";s:5:\"field\";s:4:\"slug\";s:5:\"terms\";a:2:{i:0;s:12:\"main-sub-top\";i:1;s:8:\"main-top\";}s:8:\"operator\";s:6:\"NOT IN\";}}}"},
		"shortcode_args":   {"a:154:{s:5:\"title\";s:0:\"\";s:11:\"loop_layout\";s:5:\"20661\";s:9:\"fullwidth\";s:4:\"list\";s:15:\"show_pagination\";s:2:\"on\";s:15:\"page_navigation\";s:6:\"scroll\";s:10:\"more_label\";s:10:\"더 보기\";s:7:\"columns\";s:1:\"0\";s:9:\"new_query\";s:2:\"on\";s:9:\"post_type\";s:4:\"post\";s:12:\"posts_number\";s:2:\"10\";s:13:\"offset_number\";s:1:\"3\";s:11:\"include_tax\";s:0:\"\";s:17:\"include_tax_terms\";s:0:\"\";s:11:\"exclude_tax\";s:13:\"post_curation\";s:17:\"exclude_tax_terms\";s:21:\"main-sub-top,main-top\";s:12:\"include_meta\";s:0:\"\";s:18:\"include_meta_terms\";s:0:\"\";s:21:\"include_meta_operator\";s:0:\"\";s:8:\"order_by\";s:4:\"date\";s:5:\"order\";s:4:\"desc\";s:15:\"hide_if_no_data\";s:3:\"off\";s:11:\"admin_label\";s:0:\"\";s:9:\"module_id\";s:0:\"\";s:12:\"module_class\";s:0:\"\";s:16:\"_builder_version\";s:6:\"3.0.92\";s:9:\"text_font\";s:0:\"\";s:15:\"text_text_align\";s:0:\"\";s:14:\"text_font_size\";s:4:\"14px\";s:21:\"text_font_size_tablet\";s:0:\"\";s:20:\"text_font_size_phone\";s:0:\"\";s:26:\"text_font_size_last_edited\";s:0:\"\";s:15:\"text_text_color\";s:0:\"\";s:19:\"text_letter_spacing\";s:3:\"0px\";s:26:\"text_letter_spacing_tablet\";s:0:\"\";s:25:\"text_letter_spacing_phone\";s:0:\"\";s:31:\"text_letter_spacing_last_edited\";s:0:\"\";s:16:\"text_line_height\";s:5:\"1.5em\";s:23:\"text_line_height_tablet\";s:0:\"\";s:22:\"text_line_height_phone\";s:0:\"\";s:28:\"text_line_height_last_edited\";s:0:\"\";s:22:\"text_text_shadow_style\";s:4:\"none\";s:34:\"text_text_shadow_horizontal_length\";s:3:\"0em\";s:32:\"text_text_shadow_vertical_length\";s:3:\"0em\";s:30:\"text_text_shadow_blur_strength\";s:3:\"0em\";s:22:\"text_text_shadow_color\";s:15:\"rgba(0,0,0,0.4)\";s:13:\"headings_font\";s:0:\"\";s:19:\"headings_text_align\";s:0:\"\";s:18:\"headings_font_size\";s:4:\"30px\";s:25:\"headings_font_size_tablet\";s:0:\"\";s:24:\"headings_font_size_phone\";s:0:\"\";s:30:\"headings_font_size_last_edited\";s:0:\"\";s:19:\"headings_text_color\";s:0:\"\";s:23:\"headings_letter_spacing\";s:3:\"0px\";s:30:\"headings_letter_spacing_tablet\";s:0:\"\";s:29:\"headings_letter_spacing_phone\";s:0:\"\";s:35:\"headings_letter_spacing_last_edited\";s:0:\"\";s:20:\"headings_line_height\";s:5:\"1.5em\";s:27:\"headings_line_height_tablet\";s:0:\"\";s:26:\"headings_line_height_phone\";s:0:\"\";s:32:\"headings_line_height_last_edited\";s:0:\"\";s:26:\"headings_text_shadow_style\";s:4:\"none\";s:38:\"headings_text_shadow_horizontal_length\";s:3:\"0em\";s:36:\"headings_text_shadow_vertical_length\";s:3:\"0em\";s:34:\"headings_text_shadow_blur_strength\";s:3:\"0em\";s:26:\"headings_text_shadow_color\";s:15:\"rgba(0,0,0,0.4)\";s:16:\"background_color\";s:0:\"\";s:29:\"use_background_color_gradient\";s:3:\"off\";s:31:\"background_color_gradient_start\";s:0:\"\";s:29:\"background_color_gradient_end\";s:0:\"\";s:30:\"background_color_gradient_type\";s:6:\"linear\";s:35:\"background_color_gradient_direction\";s:3:\"1px\";s:42:\"background_color_gradient_direction_radial\";s:6:\"center\";s:40:\"background_color_gradient_start_position\";s:0:\"\";s:38:\"background_color_gradient_end_position\";s:0:\"\";s:40:\"background_color_gradient_overlays_image\";s:3:\"off\";s:16:\"background_image\";s:0:\"\";s:8:\"parallax\";s:3:\"off\";s:15:\"parallax_method\";s:2:\"on\";s:15:\"background_size\";s:5:\"cover\";s:19:\"background_position\";s:6:\"center\";s:17:\"background_repeat\";s:9:\"no-repeat\";s:16:\"background_blend\";s:6:\"normal\";s:20:\"background_video_mp4\";s:0:\"\";s:21:\"background_video_webm\";s:0:\"\";s:22:\"background_video_width\";s:0:\"\";s:23:\"background_video_height\";s:0:\"\";s:18:\"allow_player_pause\";s:3:\"off\";s:18:\"__video_background\";s:0:\"\";s:12:\"border_radii\";s:6:\"on||||\";s:16:\"border_width_all\";s:0:\"\";s:16:\"border_color_all\";s:0:\"\";s:16:\"border_style_all\";s:0:\"\";s:16:\"border_width_top\";s:0:\"\";s:16:\"border_color_top\";s:0:\"\";s:16:\"border_style_top\";s:0:\"\";s:18:\"border_width_right\";s:0:\"\";s:18:\"border_color_right\";s:0:\"\";s:18:\"border_style_right\";s:0:\"\";s:19:\"border_width_bottom\";s:0:\"\";s:19:\"border_color_bottom\";s:0:\"\";s:19:\"border_style_bottom\";s:0:\"\";s:17:\"border_width_left\";s:0:\"\";s:17:\"border_color_left\";s:0:\"\";s:17:\"border_style_left\";s:0:\"\";s:13:\"custom_margin\";s:0:\"\";s:20:\"custom_margin_tablet\";s:0:\"\";s:19:\"custom_margin_phone\";s:0:\"\";s:25:\"custom_margin_last_edited\";s:0:\"\";s:21:\"padding_1_last_edited\";s:0:\"\";s:21:\"padding_2_last_edited\";s:0:\"\";s:21:\"padding_3_last_edited\";s:0:\"\";s:21:\"padding_4_last_edited\";s:0:\"\";s:14:\"custom_padding\";s:0:\"\";s:21:\"custom_padding_tablet\";s:0:\"\";s:20:\"custom_padding_phone\";s:0:\"\";s:26:\"custom_padding_last_edited\";s:0:\"\";s:15:\"animation_style\";s:4:\"none\";s:16:\"animation_repeat\";s:4:\"once\";s:19:\"animation_direction\";s:6:\"center\";s:18:\"animation_duration\";s:6:\"1000ms\";s:15:\"animation_delay\";s:3:\"0ms\";s:25:\"animation_intensity_slide\";s:3:\"50%\";s:24:\"animation_intensity_zoom\";s:3:\"50%\";s:24:\"animation_intensity_flip\";s:3:\"50%\";s:24:\"animation_intensity_fold\";s:3:\"50%\";s:24:\"animation_intensity_roll\";s:3:\"50%\";s:26:\"animation_starting_opacity\";s:2:\"0%\";s:21:\"animation_speed_curve\";s:11:\"ease-in-out\";s:16:\"box_shadow_style\";s:4:\"none\";s:21:\"box_shadow_horizontal\";s:0:\"\";s:19:\"box_shadow_vertical\";s:0:\"\";s:15:\"box_shadow_blur\";s:0:\"\";s:17:\"box_shadow_spread\";s:0:\"\";s:16:\"box_shadow_color\";s:15:\"rgba(0,0,0,0.3)\";s:19:\"box_shadow_position\";s:5:\"outer\";s:17:\"text_shadow_style\";s:4:\"none\";s:29:\"text_shadow_horizontal_length\";s:3:\"0em\";s:27:\"text_shadow_vertical_length\";s:3:\"0em\";s:25:\"text_shadow_blur_strength\";s:3:\"0em\";s:17:\"text_shadow_color\";s:15:\"rgba(0,0,0,0.4)\";s:17:\"custom_css_before\";s:0:\"\";s:23:\"custom_css_main_element\";s:0:\"\";s:16:\"custom_css_after\";s:0:\"\";s:8:\"disabled\";s:3:\"off\";s:11:\"disabled_on\";s:0:\"\";s:13:\"global_module\";s:0:\"\";s:10:\"saved_tabs\";s:0:\"\";s:10:\"ab_subject\";s:0:\"\";s:13:\"ab_subject_id\";s:0:\"\";s:7:\"ab_goal\";s:0:\"\";s:6:\"locked\";s:0:\"\";s:13:\"template_type\";s:0:\"\";s:12:\"inline_fonts\";s:0:\"\";s:9:\"collapsed\";s:0:\"\";}"},
		"paged":            {strconv.Itoa(page)},
		"total_posts":      {"1375"},
		"posts_per_page":   {"10"},
		"page_navigation":  {"scroll"},
		"module_class":     {"et_pb_cpt_loop_archive_6 et_pb_cpt_archive_list"}})
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		src, ok := s.Attr("src")
		if !ok {
			fmt.Printf("not fine")
			return
		}
		fmt.Printf("iimg %d: %s\n", i, src)
	})
}
func Coindesk_korea_scrape_range(start_page int, end_page int) {
	for i := start_page; i < end_page; i++ {
		Coindesk_korea_scrape_page(i)
	}
}
