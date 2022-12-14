package mtkJpegEnhanceSkia
import (
	"android/soong/android"
	"android/soong/cc"
	"android/soong/android/mediatek"
	//"github.com/google/blueprint"
)

func mtkJpegEnhanceSkiaDefaults(ctx android.LoadHookContext) {
	type props struct {
		Target struct {
			Android struct {
				Cflags []string
				Local_include_dirs []string
				Include_dirs []string
				Srcs []string
				Shared_libs []string
			}
		}
	}
	p := &props{}
	//featureValue := android.MtkFeatureValues

	var localIncludeDirs []string
	var includeDirs []string
	var srcs []string
	var sharedlibs []string

	if mediatek.GetFeature("MTK_JPEG_HW_RESIZER_TYPE") == "HW_RESIZER_TYPE_2"  {
			//graphics flags
			p.Target.Android.Cflags = append(p.Target.Android.Cflags, "-D__MTK_TRACE_SKIA__")
			p.Target.Android.Cflags = append(p.Target.Android.Cflags, "-D__MTK_TRACE_MT_BLITTER__")
			p.Target.Android.Cflags = append(p.Target.Android.Cflags, "-D__MTK_AFFINE_TRANSFORM_EXT__")
			p.Target.Android.Cflags = append(p.Target.Android.Cflags, "-D__MULTI_THREADS_OPTIMIZE_2D__")
			//p.Target.Android.Cflags = append(p.Target.Android.Cflags, "-D__MTK_DUMP_DRAW_BITMAP__")

			//jpeg flags
			//p.Target.Android.Cflags = append(p.Target.Android.Cflags, "-DMTK_JPEG_HW_DECODER")
			p.Target.Android.Cflags = append(p.Target.Android.Cflags, "-DMTK_JPEG_HW_REGION_RESIZER")
			p.Target.Android.Cflags = append(p.Target.Android.Cflags, "-DMTK_SKIA_MULTI_THREAD_JPEG_REGION")
			p.Target.Android.Cflags = append(p.Target.Android.Cflags, "-DSK_JPEG_INDEX_SUPPORTED_MTK")
			p.Target.Android.Cflags = append(p.Target.Android.Cflags, "-DMTK_IMAGE_ENABLE_PQ_FOR_JPEG")

			//graphics local dirs
			localIncludeDirs = append(localIncludeDirs, "mtk_opt/include/adapter")
			localIncludeDirs = append(localIncludeDirs, "mtk_opt/src/adapter")

			//jpeg local dirs
			localIncludeDirs = append(localIncludeDirs, "include/mtk/")

			//jpeg non-local dirs
			includeDirs = append(includeDirs, "system/core/libion/include/")		//ion/ion.h
			includeDirs = append(includeDirs, "system/core/include/utils/")
			includeDirs = append(includeDirs, "vendor/mediatek/proprietary/external/")	//libion_mtk/include/ion.h
			includeDirs = append(includeDirs, "device/mediatek/common/kernel-headers/")		//linux/mtk_ion.h
			includeDirs = append(includeDirs, "vendor/mediatek/proprietary/external/libjpeg-alpha/include/")
			includeDirs = append(includeDirs, "vendor/mediatek/proprietary/hardware/dpframework/include/")

			//graphics srcs
			srcs = append(srcs, "mtk_opt/src/adapter/SkBlitterAdapterHandler.cpp")
			srcs = append(srcs, "mtk_opt/src/adapter/SkBlitterMTAdapter.cpp")

			//jpeg srcs
			srcs = append(srcs, "src/mtk/SkError.cpp")
			srcs = append(srcs, "src/mtk/SkImageDecoder.cpp")
			srcs = append(srcs, "src/mtk/SkImageDecoder_FactoryDefault.cpp")
			srcs = append(srcs, "src/mtk/SkImageDecoder_FactoryRegistrar.cpp")
			srcs = append(srcs, "src/mtk/SkImageDecoder_libjpeg.cpp")
			srcs = append(srcs, "src/mtk/SkJpegCodec.cpp")
			srcs = append(srcs, "src/mtk/SkJpegDecoderMgr.cpp")
			srcs = append(srcs, "src/mtk/SkJpegUtility.cpp")
			srcs = append(srcs, "src/mtk/SkScaledBitmapSampler.cpp")

			//jpeg legacy shared libs
			sharedlibs = append(sharedlibs, "libion")
			//sharedlibs = append(sharedlibs, "libmhalImageCodec_sys")
			sharedlibs = append(sharedlibs, "libdpframework_mtk")
			sharedlibs = append(sharedlibs, "libion_mtk_sys")
			sharedlibs = append(sharedlibs, "libjpeg-alpha")

	} else {
			srcs = append(srcs, "src/codec/SkJpegCodec.cpp")
	}

	p.Target.Android.Local_include_dirs = localIncludeDirs
	p.Target.Android.Include_dirs = includeDirs
	p.Target.Android.Srcs = srcs
	p.Target.Android.Shared_libs = sharedlibs

	ctx.AppendProperties(p)
}

func init() {
	android.RegisterModuleType("mtk_jpeg_enhance_skia_defaults", mtkJpegEnhanceSkiaDefaultsFactory)
}

func mtkJpegEnhanceSkiaDefaultsFactory() (android.Module) {
	module := cc.DefaultsFactory()
	android.AddLoadHook(module, mtkJpegEnhanceSkiaDefaults)
	return module
}