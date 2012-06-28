/**
 * Copyright 2002-2012 Evgeny Gryaznov
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package org.textway.lapg.test.eval;

import org.junit.Test;
import org.textway.lapg.test.gen.LapgTemplatesTestHelper;

import java.util.HashMap;
import java.util.Map;

/**
 * Gryaznov Evgeny, 2/26/12
 */
public class GenericParserTest {

	@Test
	public void testGenericParser() {
		new LapgTemplatesTestHelperEx().gentest(
				"java.main", "tests/org/textway/lapg/test/eval/templates",
				"../lapg-core/src/org/textway/lapg/eval",
				new String[]{"GenericParser.java", "GenericLexer.java"});
	}

	private static class LapgTemplatesTestHelperEx extends LapgTemplatesTestHelper {
		@Override
		protected Map<String, Object> createOptions() {
			HashMap<String, Object> res = new HashMap<String, Object>();
			res.put("prefix", "Generic");
			res.put("package", "org.textway.lapg.eval");
			res.put("positions", "line,offset");
			res.put("endpositions", "offset");
			res.put("lexerInput", "buffered");
			res.put("maxtoken", 2048);
			res.put("stack", 1024);
			res.put("genast", false);
			res.put("gentree", false);
			res.put("genCleanup", true);
			return res;
		}
	}

}